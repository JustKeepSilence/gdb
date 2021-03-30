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
	"github.com/gin-contrib/cors"
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
	"strconv"
	"strings"
	"time"
)

// web server

func portInUse(portNumber int64) (int, error) {
	res := -1
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
	return res, nil
}

var (
	g errgroup.Group
)

// web app router
func appRouter(g *Gdb, authorization, logWriting bool, level logLevel) http.Handler {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			time.Sleep(60 * time.Second)
		}
	}()
	router := gin.New()
	pprof.Register(router)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	})) // customer console writing,disable console writing
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if _, ok := recovered.(string); ok {
			if c.Request.Method == "GET" {
				//_ = g.writeLog(Error, "get", err, c.Request.URL.String(), c.Request.URL.String())
			} else if c.Request.Method == "POST" {
				// b, _ := ioutil.ReadAll(c.Request.Body)
				//_ = g.writeLog(Error, "post", fmt.Sprintf("%s", b), c.Request.URL.String(), c.Request.URL.String())
			}
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Authorization", ""},
	}))
	router.Use(cors.Default()) // allow all cors
	if authorization {
		router.Use(g.authorization()) // authorization
	}
	if logWriting {
		router.Use(g.setLogHeader(level))
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
		item.POST("/getItems", g.getItemsHandler)
		item.POST("/getItemsWithCount", g.handleGetItemsWithCount) // get item with  count
		item.POST("/updateItems", g.updateItemsHandler)
	}
	data := router.Group("/data") // data handler
	{
		data.POST("/batchWrite", g.batchWriteHandler)
		data.POST("/getRealTimeData", g.getRealTimeDataHandler)
		data.POST("/getHistoricalData", g.getHistoricalDataHandler)
		data.POST("/getHistoricalDataWithStamp", g.getHistoricalDataWithStampHandler)
		data.POST("/getHistoricalDataWithCondition", g.getHistoricalDataWithConditionHandler)
		data.POST("/getDbInfo", g.getDbInfoHandler)
		data.POST("/getDbSpeedHistory", g.getDbSpeedHistoryHandler)
		data.POST("/getRawData", g.getRawDataHandler)
	}
	pageRequest := router.Group("/page") // page request handler
	{
		pageRequest.POST("/userLogin", g.handleUserLogin) // user login
		pageRequest.GET("/userLogout/:userName", g.handleUserLogout)
		pageRequest.POST("/getUserInfo", g.handleGetUerInfo)          // get user info
		pageRequest.POST("/uploadFile", g.handleUploadFile)           // upload file
		pageRequest.POST("/addItemsByExcel", g.handleAddItemsByExcel) // add item by excel
		pageRequest.GET("/getJsCode/:fileName", g.getJsCodeHandler)   // get js code
		pageRequest.GET("/getLogs", g.getLogsHandler)                 // get logs
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
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/groups", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.GET("/calc", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.GET("/document", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.GET("/userManagement", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.GET("/log", func(c *gin.Context) {
		c.Request.URL.Path = "/index"
		router.HandleContext(c)
	})
	router.Static("/static", "./dist/static") // load static files
	router.LoadHTMLGlob("./*.html")           // render html template
	return router
}

// initial gdb service according to configs.json file
func StartDbServer(configs Config) error {
	dbPath, itemDbPath, port, ip, mode, ca, caCertificateName, serverCertificateName, selfSignedCa :=
		configs.DbPath, configs.ItemDbPath, configs.Port, configs.IP, configs.Mode,
		configs.Ca, configs.CaCertificateName, configs.ServerCertificateName, configs.SelfSignedCa
	if mode == "" {
		mode = "http"
	}
	checkResult, err := portInUse(port)
	if err != nil {
		return fmt.Errorf("%s: fail in checking port %d: %s", time.Now().Format(timeFormatString), port, err)
	}
	if checkResult != -1 {
		// used
		return fmt.Errorf("%s: Failed to start web service: Port number %d is already occupied, process PID is %d, please consider using taskkill /f /pid %d to terminate the process", time.Now().Format("2006-01-02 15:04:05"), port, checkResult, checkResult)
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
	if mode == "https" && ca {
		// use ca root
		if certificate, err := tls.LoadX509KeyPair("./ssl/"+serverCertificateName+".crt", "./ssl/"+serverCertificateName+".key"); err != nil {
			return err
		} else {
			if caFile, err := ioutil.ReadFile("./ssl/" + caCertificateName + ".crt"); err != nil {
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
				cred := credentials.NewTLS(&tls.Config{
					Certificates: []tls.Certificate{certificate},
					ClientAuth:   tls.RequireAndVerifyClientCert,
					ClientCAs:    certPool,
				})
				s = grpc.NewServer(grpc.ChainUnaryInterceptor(se.panicInterceptor, se.authInterceptor, se.logInterceptor),
					grpc.ChainStreamInterceptor(se.panicWithServerStreamInterceptor, se.authWithServerStreamInterceptor),
					grpc.Creds(cred))
			}
		}
	} else {
		s = grpc.NewServer(grpc.ChainUnaryInterceptor(se.panicInterceptor, se.authInterceptor, se.logInterceptor),
			grpc.ChainStreamInterceptor(se.panicWithServerStreamInterceptor, se.authWithServerStreamInterceptor))
	}
	pb.RegisterGroupServer(s, se)
	pb.RegisterItemServer(s, se)
	pb.RegisterDataServer(s, se)
	pb.RegisterPageServer(s, se)
	pb.RegisterCalcServer(s, se)
	fmt.Printf("%s: launch gdb service successfully!: %s \n", time.Now().Format(timeFormatString), address)
	if mode == "http" {
		h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				s.ServeHTTP(w, r)
			} else {
				appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level)
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
			if err := http.ListenAndServeTLS(address, "./ssl/"+serverCertificateName+".crt", "./ssl/"+serverCertificateName+".key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
					s.ServeHTTP(w, r)
				} else {
					appRouter(gdb, configs.Authorization, configs.LogWriting, configs.Level)
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
