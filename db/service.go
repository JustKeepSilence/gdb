/*
creatTime: 2021/2/7
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "github.com/JustKeepSilence/gdb/model"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// web server

func portInUse(portNumber int64) (int, error) {
	res := -1
	if runtime.GOOS == "windows" {
		cmdStr := fmt.Sprintf(`netstat -ano -p tcp | findstr %d`, portNumber)
		outBytes, _ := exec.Command("cmd", "/C", cmdStr, " ").Output()
		resStr := fmt.Sprintf("%s", outBytes)
		r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
		//  TCP    192.168.0.199:8082     0.0.0.0:0              LISTENING       9404
		if len(r) > 0 {
			pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
			if err != nil {
				res = -1
			} else {
				res = pid
			}
		}
	} else {
		// linux
		cmdStr := fmt.Sprintf("netstat -anp |grep %d", portNumber)
		outBytes, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
		resStr := fmt.Sprintf("%s", outBytes)
		if resStr == "" {
			return -1, nil
		} else {
			r := regexp.MustCompile(`\s\d+/`).FindAllString(resStr, -1)
			if len(r) > 0 {
				if pid, err := strconv.Atoi(strings.TrimSpace(strings.Replace(r[0], "/", "", -1))); err != nil {
					return -1, err
				} else {
					return pid, nil
				}
			}
		}
	}
	return res, nil
}

func checkConfigs(mode, level string) (string, bool) {
	if mode != "http" && mode != "https" && mode != "gRPC" {
		return "invalid mode " + mode + ", mode must be http, https or gRPC", false
	} else if level != "Info" && level != "Error" {
		return "invalid level " + level + ", level must be Info or Error", false
	} else {
		return "", true
	}
}

var (
	g errgroup.Group
)

// web app router
func appRouter(g *Gdb, authorization, logWriting bool, level string) http.Handler {
	router := gin.New()
	pprof.Register(router)
	//router.Use(g.allowOptionRequest())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	})) // customer console writing,disable console writing
	router.Use(g.corsMiddleware())
	if authorization {
		router.Use(g.authorizationMiddleware()) // authorization
	}
	if logWriting {
		router.Use(g.setLogHeaderMiddleware(level))
	}
	group := router.Group("/group") // group handler
	{
		group.POST("/addGroups", g.addGroupsHandler)               // add
		group.POST("/deleteGroups", g.deleteGroupsHandler)         // delete
		group.POST("/getGroups", g.getGroupsHandler)               // get
		group.POST("/getGroupProperty", g.getGroupPropertyHandler) // update
		group.POST("/updateGroupNames", g.updateGroupNamesHandler) // update
		group.POST("/updateGroupColumnNames", g.updateGroupColumnNamesHandler)
		group.POST("/deleteGroupColumns", g.deleteGroupColumnsHandler)
		group.POST("/addGroupColumns", g.addGroupColumnsHandler)
		group.POST("/cleanGroupItems", g.cleanGroupItemsHandler)
	}
	item := router.Group("/item") // item handler
	{
		item.POST("/addItems", g.addItemsHandler)
		item.POST("/deleteItems", g.deleteItemsHandler)
		item.POST("/getItems", g.getItemsHandler)
		item.POST("/getItemsWithCount", g.handleGetItemsWithCount) // get item with  count
		item.POST("/updateItems", g.updateItemsHandler)
		item.POST("/checkItems", g.checkItemsHandler)
	}
	data := router.Group("/data") // data handler
	{
		data.POST("/batchWrite", g.batchWriteHandler)
		data.POST("/batchWriteHistoricalData", g.batchWriteHistoricalDataHandler)
		data.POST("/getRealTimeData", g.getRealTimeDataHandler)
		data.POST("/getHistoricalData", g.getHistoricalDataHandler)
		data.POST("/getHistoricalDataWithStamp", g.getHistoricalDataWithStampHandler)
		data.POST("/getHistoricalDataWithCondition", g.getHistoricalDataWithConditionHandler)
		data.POST("/getDbInfo", g.getDbInfoHandler)
		data.POST("/getDbInfoHistory", g.getDbInfoHistoryHandler)
		data.POST("/getRawData", g.getRawDataHandler)
	}
	pageRequest := router.Group("/page") // page request handler
	{
		pageRequest.POST("/userLogin", g.handleUserLogin) // user login
		pageRequest.POST("/userLogOut", g.handleUserLogout)
		pageRequest.POST("/getUserInfo", g.getUerInfoHandler) // get user info
		pageRequest.POST("/getUsers", g.getUsersHandler)
		pageRequest.POST("/addUsers", g.addUsersHandler)
		pageRequest.POST("/deleteUsers", g.deleteUsersHandler)
		pageRequest.POST("/updateUsers", g.updateUsersHandler)
		pageRequest.POST("/uploadFile", g.handleUploadFile)                  // upload file
		pageRequest.POST("/httpsUploadFile", g.handleHttpsUploadFile)        // https upload file
		pageRequest.POST("/addItemsByExcel", g.handleAddItemsByExcelHandler) // add item by excel
		pageRequest.POST("/importHistoryByExcel", g.importHistoryByExcelHandler)
		pageRequest.POST("/getJsCode", g.getJsCodeHandler) // get js code
		pageRequest.POST("/getLogs", g.getLogsHandler)     // get logs
		pageRequest.POST("/deleteLogs", g.deleteLogsHandler)
		pageRequest.POST("/downloadFile", g.downloadFileHandler)
	}
	calcRequest := router.Group("/calculation")
	{
		calcRequest.POST("/addCalcItem", g.addCalcItemHandler) // add calc item
		calcRequest.POST("/getCalcItems", g.getCalcItemsHandler)
		calcRequest.POST("/updateCalcItem", g.updateCalcItemHandler)
		calcRequest.POST("/startCalcItem", g.startCalculationItemHandler)
		calcRequest.POST("/stopCalcItem", g.stopCalculationItemHandler)
		calcRequest.POST("/deleteCalcItem", g.deleteCalculationItemHandler)
	}
	// web page handler
	return router
}

// StartDbServer initial gdb service according to configs.json file
// Both http and https modes support gRPC. If it is https mode,
//it supports gRPC mode with CA certificate and without CA certificate enabled.
//Self-visa certificates are only useful on linux and mac
func StartDbServer(configs Config) error {
	dbPath, itemDbPath, port, ip, mode, ca, caCertificateName, serverCertificateName, serverKeyName, selfSignedCa :=
		configs.DbPath, configs.ItemDbPath, configs.Port, configs.IP, configs.Mode,
		configs.Ca, configs.CaCertificateName, configs.ServerCertificateName, configs.ServerKeyName, configs.SelfSignedCa
	if mode == "" {
		mode = "http"
	}
	if result, ok := checkConfigs(mode, configs.Level); !ok {
		return fmt.Errorf(result)
	}
	checkResult, err := portInUse(port)
	if err != nil {
		return fmt.Errorf("%s: fail in checking port %d: %s", time.Now().Format(timeFormatString), port, err)
	}
	if checkResult != -1 {
		// used
		if runtime.GOOS == "windows" {
			return fmt.Errorf("%s: Failed to start web service: Port number %d is already occupied, process PID is %d, please consider using taskkill /f /pid %d to terminate the process", time.Now().Format("2006-01-02 15:04:05"), port, checkResult, checkResult)
		} else {
			return fmt.Errorf("%s: Failed to start web service: Port number %d is already occupied, process PID is %d, please consider using kill -9 /pid %d to terminate the process", time.Now().Format("2006-01-02 15:04:05"), port, checkResult, checkResult)
		}
	}
	if len(ip) == 0 {
		// not config ip
		ip = getLocalIp()
	}
	gin.SetMode(gin.ReleaseMode)                  // production
	address := ip + ":" + fmt.Sprintf("%d", port) // base url of web server
	gdb, err := NewGdb(dbPath, itemDbPath)
	if err != nil {
		return err
	}
	se := &server{gdb: gdb, configs: configs}
	var s *grpc.Server
	if mode == "https" {
		// https mode and use ca root
		var cred credentials.TransportCredentials
		if ca {
			if certificate, err := tls.LoadX509KeyPair("./ssl/"+serverCertificateName, "./ssl/"+serverKeyName); err != nil {
				return err
			} else {
				if caFile, err := ioutil.ReadFile("./ssl/" + caCertificateName); err != nil {
					return err
				} else {
					if selfSignedCa {
						// only useful in linux and macOS
						if sysPool, err := x509.SystemCertPool(); err != nil {
							return err
						} else {
							if ok := sysPool.AppendCertsFromPEM(caFile); !ok {
								return fmt.Errorf("failed to add ca to system root pool")
							} // add root ca file to system root ca pool
						}
					}
					certPool := x509.NewCertPool()
					if ok := certPool.AppendCertsFromPEM(caFile); !ok {
						return fmt.Errorf("failed to add ca to root pool")
					}
					cred = credentials.NewTLS(&tls.Config{
						Certificates: []tls.Certificate{certificate},
						ClientAuth:   tls.RequireAndVerifyClientCert,
						ClientCAs:    certPool,
					})
				}
			}
		} else {
			var err1 error
			cred, err1 = credentials.NewServerTLSFromFile("./ssl/"+serverCertificateName, "./ssl/"+serverKeyName)
			if err1 != nil {
				return err1
			}
		}
		s = grpc.NewServer(grpc.ChainUnaryInterceptor(se.panicInterceptor, se.authInterceptor, se.logInterceptor),
			grpc.ChainStreamInterceptor(se.panicWithServerStreamInterceptor, se.authWithServerStreamInterceptor),
			grpc.Creds(cred), grpc.MaxRecvMsgSize(1024*1024*100)) // 100 MB
	} else {
		// http gRPC
		s = grpc.NewServer(grpc.ChainUnaryInterceptor(se.panicInterceptor, se.authInterceptor, se.logInterceptor),
			grpc.ChainStreamInterceptor(se.panicWithServerStreamInterceptor, se.authWithServerStreamInterceptor), grpc.MaxRecvMsgSize(1024*1024*100))
	}
	// register gRPC
	pb.RegisterGroupServer(s, se)
	pb.RegisterItemServer(s, se)
	pb.RegisterDataServer(s, se)
	pb.RegisterPageServer(s, se)
	pb.RegisterCalcServer(s, se)
	fmt.Printf("%s: launch gdb service successfully!: %s, mode: %s,authorization: %s \n ", time.Now().Format(timeFormatString), address, mode, strconv.FormatBool(configs.Authorization))
	if mode == "http" {
		// http mode
		h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level).ServeHTTP(w, r)
			}
		}), &http2.Server{})
		hs := &http.Server{Addr: address, Handler: h2Handler}
		g.Go(func() error {
			if err := hs.ListenAndServe(); err != nil {
				return err
			} else {
				return nil
			}
		})
	} else {
		// https mode
		g.Go(func() error {
			if err := http.ListenAndServeTLS(address, "./ssl/"+serverCertificateName, "./ssl/"+serverKeyName, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
					s.ServeHTTP(w, r)
				} else {
					appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level).ServeHTTP(w, r)
				}
			})); err != nil && err != http.ErrServerClosed {
				return err
			} else {
				return nil
			}
		})
	}
	g.Go(gdb.getProcessInfo) // monitor
	g.Go(gdb.calc)           // calc goroutine
	if err := g.Wait(); err != nil {
		return fmt.Errorf("%s: runTimeError: %s", time.Now().Format(timeFormatString), err.Error())
	}
	return nil
}
