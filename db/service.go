// +build gdbServer

/*
creatTime: 2021/2/7
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

import (
	"context"
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
	"os"
	"os/exec"
	"os/signal"
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

// web app router
func appRouter(g *Gdb, authorization, logWriting bool, level string) http.Handler {
	router := gin.New()
	pprof.Register(router)
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
	}
	item := router.Group("/item") // item handler
	{
		item.POST("/addItems", g.addItemsHandler)
		item.POST("/deleteItems", g.deleteItemsHandler)
		item.POST("/getItemsWithCount", g.handleGetItemsWithCount) // get item with  count
		item.POST("/updateItems", g.updateItemsHandler)
		item.POST("/checkItems", g.checkItemsHandler)
		item.POST("/cleanGroupItems", g.cleanGroupItemsHandler)
	}
	data := router.Group("/data") // data handler
	{
		data.POST("/batchWriteFloatData", g.batchWriteFloatDataHandler)
		data.POST("/batchWriteIntData", g.batchWriteIntDataHandler)
		data.POST("/batchWriteBoolData", g.batchWriteBoolDataHandler)
		data.POST("/batchWriteStringData", g.batchWriteStringDataHandler)
		data.POST("/batchWriteFloatHistoricalData", g.batchWriteFloatHistoricalDataHandler)
		data.POST("/batchWriteIntHistoricalData", g.batchWriteIntHistoricalDataHandler)
		data.POST("/batchWriteStringHistoricalData", g.batchWriteStringHistoricalDataHandler)
		data.POST("/batchWriteBoolHistoricalData", g.batchWriteBoolHistoricalDataHandler)
		data.POST("/getRealTimeData", g.getRealTimeDataHandler)
		data.POST("/getFloatHistoricalData", g.getFloatHistoricalDataHandler)
		data.POST("/getIntHistoricalData", g.getIntHistoricalDataHandler)
		data.POST("/getStringHistoricalData", g.getStringHistoricalDataHandler)
		data.POST("/getBoolHistoricalData", g.getBoolHistoricalDataHandler)
		data.POST("/getFloatRawHistoricalData", g.getFloatRawHistoricalDataHandler)
		data.POST("/getIntRawHistoricalData", g.getIntRawHistoricalDataHandler)
		data.POST("/getStringRawHistoricalData", g.getStringRawHistoricalDataHandler)
		data.POST("/getBoolRawHistoricalData", g.getBoolRawHistoricalDataHandler)
		data.POST("/getFloatHistoricalDataWithStamp", g.getFloatHistoricalDataWithStampHandler)
		data.POST("/getIntHistoricalDataWithStamp", g.getIntHistoricalDataWithStampHandler)
		data.POST("/getStringHistoricalDataWithStamp", g.getStringHistoricalDataWithStampHandler)
		data.POST("/getBoolHistoricalDataWithStamp", g.getBoolHistoricalDataWithStampHandler)
		data.POST("/getFloatHistoricalDataWithCondition", g.getFloatHistoricalDataWithConditionHandler)
		data.POST("/getIntHistoricalDataWithCondition", g.getIntHistoricalDataWithConditionHandler)
		data.POST("/getStringHistoricalDataWithCondition", g.getStringHistoricalDataWithConditionHandler)
		data.POST("/getBoolHistoricalDataWithCondition", g.getBoolHistoricalDataWithConditionHandler)
		data.POST("/deleteFloatHistoricalData", g.deleteFloatHistoricalDataHandler)
		data.POST("/deleteIntHistoricalData", g.deleteIntHistoricalDataHandler)
		data.POST("/deleteStringHistoricalData", g.deleteStringHistoricalDataHandler)
		data.POST("/deleteBoolHistoricalData", g.deleteBoolHistoricalDataHandler)
		data.POST("/cleanItemData", g.cleanItemDataHandler)
		data.POST("/reLoadDb", g.reLoadDbHandler)
	}
	page := router.Group("/page") // page request handler
	{
		page.POST("/userLogin", g.handleUserLogin) // user login
		page.POST("/userLogOut", g.handleUserLogout)
		page.POST("/getUserInfo", g.getUerInfoHandler) // get user info
		page.POST("/getUsers", g.getUsersHandler)
		page.POST("/addUsers", g.addUsersHandler)
		page.POST("/deleteUsers", g.deleteUsersHandler)
		page.POST("/updateUsers", g.updateUsersHandler)
		page.POST("/uploadFile", g.handleUploadFile)                  // upload file
		page.POST("/httpsUploadFile", g.handleHttpsUploadFile)        // https upload file
		page.POST("/addItemsByExcel", g.handleAddItemsByExcelHandler) // add item by excel
		page.POST("/importHistoryByExcel", g.importHistoryByExcelHandler)
		page.POST("/getJsCode", g.getJsCodeHandler) // get js code
		page.POST("/getLogs", g.getLogsHandler)     // get logs
		page.POST("/deleteLogs", g.deleteLogsHandler)
		page.POST("/downloadFile", g.downloadFileHandler)
		page.POST("/getDbSize", g.getDbSizeHandler)
		page.POST("/getDbInfo", g.getDbInfoHandler)
		page.POST("/getDbInfoHistory", g.getDbInfoHistoryHandler)
		page.POST("/getRoutes", g.getRoutesHandler)
		page.POST("/deleteRoutes", g.deleteRoutesHandler)
		page.POST("/addRoutes", g.addRoutesHandler)
		page.POST("/addUserRoutes", g.addUserRoutesHandler)
		page.POST("/deleteUserRoutes", g.deleteUserRoutesHandler)
		page.POST("/getAllRoutes", g.getAllRoutesHandler)
		page.POST("/checkRoutes", g.checkRoutesHandler)
	}
	calc := router.Group("/calculation")
	{
		calc.POST("/testCalcItem", g.testCalcItemHandler)
		calc.POST("/addCalcItem", g.addCalcItemHandler) // add calc item
		calc.POST("/getCalcItems", g.getCalcItemsHandler)
		calc.POST("/updateCalcItem", g.updateCalcItemHandler)
		calc.POST("/startCalcItem", g.startCalculationItemHandler)
		calc.POST("/stopCalcItem", g.stopCalculationItemHandler)
		calc.POST("/deleteCalcItem", g.deleteCalculationItemHandler)
	}
	cmd := router.Group("/cmd")
	{
		cmd.GET("/getCmdInfo/:name", g.getCmdInfoHandler)
	}
	// web page handler
	return router
}

// StartDbServer initial gdb service according to configs.json file
// Both http and https modes support gRPC. If it is https mode,
//it supports gRPC mode with CA certificate and without CA certificate enabled.
//Self-visa certificates are only useful on linux and mac
func StartDbServer(configs Config) error {
	g := errgroup.Group{}
	dbPath, port, ip, mode, ca, caCertificateName, serverCertificateName, serverKeyName, selfSignedCa, userRedis, redisIP, redisPort, redisPassWord, redisDb, keyName, driverName, dsn, rtTimeDuration, hisTimeDuration :=
		configs.DbPath, configs.Port, configs.IP, configs.Mode, configs.Ca, configs.CaCertificateName, configs.ServerCertificateName, configs.ServerKeyName, configs.SelfSignedCa, configs.UseRedis, configs.RedisIp, configs.RedisPort,
		configs.RedisPassWord, configs.RedisDb, configs.KeyName, configs.DriverName, configs.Dsn, configs.RtTimeDuration, configs.HisTimeDuration
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
	gdb, err := innerNewGdb(dbPath, userRedis, redisIP, redisPort, redisPassWord, redisDb, keyName, driverName, dsn, rtTimeDuration, hisTimeDuration)
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
			grpc.Creds(cred), grpc.MaxRecvMsgSize(1024*1024*1024))
	} else {
		// http gRPC, for http gRPC ,th max message receive is 1G
		s = grpc.NewServer(grpc.ChainUnaryInterceptor(se.panicInterceptor, se.authInterceptor, se.logInterceptor),
			grpc.ChainStreamInterceptor(se.panicWithServerStreamInterceptor, se.authWithServerStreamInterceptor), grpc.MaxRecvMsgSize(1024*1024*1024))
	}
	// register gRPC
	pb.RegisterGroupServer(s, se)
	pb.RegisterItemServer(s, se)
	pb.RegisterDataServer(s, se)
	pb.RegisterPageServer(s, se)
	pb.RegisterCalcServer(s, se)
	fmt.Printf("%s: launch gdb service successfully!: %s, mode: %s,authorization: %s, you should use Ctrl+C to stop gdbService \n ", time.Now().Format(timeFormatString), address, mode, strconv.FormatBool(configs.Authorization))
	var hs *http.Server
	if mode == "http" {
		// http mode
		h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level).ServeHTTP(w, r)
			}
		}), &http2.Server{})
		hs = &http.Server{Addr: address, Handler: h2Handler}
		g.Go(func() error {
			if err := hs.ListenAndServe(); err != nil {
				return err
			} else {
				return nil
			}
		})
	} else {
		// https mode
		hs = &http.Server{Addr: address, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level).ServeHTTP(w, r)
			}
		})}
		g.Go(func() error {
			if err := hs.ListenAndServeTLS("./ssl/"+serverCertificateName, "./ssl/"+serverKeyName); err != nil {
				return err
			} else {
				return nil
			}
		})
	}
	g.Go(func() error {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		<-ctx.Done()
		stop()
		fmt.Println(time.Now().Format(timeFormatString), ": system is stopping...")
		timeoutCxt, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// close tcp
		if err := hs.Shutdown(timeoutCxt); err != nil {
			return err
		}
		// close gRPC
		s.GracefulStop()
		_ = gdb.CloseGdb()
		fmt.Println(time.Now().Format(timeFormatString), ": system stop successfully, will exit in 3 seconds...")
		time.Sleep(3 * time.Second)
		os.Exit(-1)
		return nil
	})
	g.Go(gdb.getProcessInfo) // monitor
	g.Go(gdb.calc)           // calc goroutine
	g.Go(gdb.syncRtData)
	g.Go(gdb.syncHisData)
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
