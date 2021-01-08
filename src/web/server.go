/*
creatTime: 2020/11/8 21:19
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
*/

package web

import (
	"db"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
	"utils"
)

// Check whether a port is occupied, the returned value is -1, which means it was not found
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
func appRouter(dbPath string, ldb *db.LevelDb) http.Handler {
	defer func() {
		if err := recover();err!=nil{
			fmt.Println(err)
			time.Sleep(60 * time.Second)
		}
	}()
	if err := ldb.InitialDb(dbPath, -1);err!=nil{
		utils.WriteError("", "", "", err.Error())
		return nil
	}
	router := gin.New()
	pprof.Register(router)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
	})) // customer console writing
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string);ok{
			if c.Request.Method == "GET"{
				utils.WriteError(c.Request.URL.String(), "GET", c.Request.URL.String(), err)
			}else if c.Request.Method == "POST"{
				b, _ := ioutil.ReadAll(c.Request.Body)
				utils.WriteError(c.Request.URL.String(), "POST", fmt.Sprintf("%s", b), err)
			}
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{"Authorization", ""},
	}))
	router.Use(cors.Default())  // allow all cors
	group := router.Group("/group")  // group handler
	{
		group.POST("/addGroups", ldb.AddGroupsHandler)  // add
		group.POST("/deleteGroups", ldb.DeleteGroupsHandler)  // delete
		group.POST("/getGroups", ldb.GetGroupsHandler)  // get
		group.POST("/getGroupProperty", ldb.GetGroupPropertyHandler) // update
		group.POST("/updateGroupNames", ldb.UpdateGroupNamesHandler)  // update
		group.POST("/updateGroupColumnNames", ldb.UpdateGroupColumnNamesHandler)
		group.POST("/deleteGroupColumns", ldb.DeleteGroupColumnsHandler)
		group.POST("/addGroupColumns", ldb.AddGroupColumnsHandler)
	}
	item := router.Group("/item")  // item handler
	{
		item.POST("/addItems", ldb.AddItemsHandler)
		item.POST("/deleteItems", ldb.DeleteItemsHandler)
		item.POST("/getItems", ldb.GetItemsHandler)
		item.POST("/updateItems", ldb.UpdateItemsHandler)
	}
	data := router.Group("/data")  // data handler
	{
		data.POST("/batchWrite", ldb.BatchWriteHandler)
		data.POST("/getRealTimeData", ldb.GetRealTimeDataHandler)
		data.POST("/getHistoricalData", ldb.GetHistoricalDataHandler)
		data.POST("/getHistoricalDataWithStamp", ldb.GetHistoricalDataWithStampHandler)
		data.POST("/getHistoricalDataWithCondition", ldb.GetHistoricalDataWithConditionHandler)
		data.POST("/getDbInfo", ldb.GetDbInfoHandler)
		data.POST("/getDbSpeedHistory", ldb.GetDbSpeedHistoryHandler)
	}
	pageRequest := router.Group("/page")  // page request handler
	{
		pageRequest.POST("/userLogin", ldb.HandleUserLogin)  // user login
		pageRequest.POST("/getUserInfo", ldb.HandleGetUerInfo)  // get user info
		pageRequest.POST("/uploadFile", ldb.HandleUploadFile)  // upload file
		pageRequest.POST("/addItemsByExcel", ldb.HandleAddItemsByExcel) // add item by excel
		pageRequest.POST("/getItemsWithCount", ldb.HandleGetItemsWithCount)  // get item with  count
		pageRequest.GET("/getJsCode/:fileName", ldb.GetJsCodeHandler)  // get js code
		pageRequest.GET("/getLogs", ldb.GetLogsHandler)  // get logs
	}
	calcRequest := router.Group("/calculation")
	{
		calcRequest.POST("/addCalcItem", ldb.AddCalcItemHandler)  // add calc item
		calcRequest.POST("/getCalcItem", ldb.GetCalcItemHandler)
		calcRequest.POST("/updateCalcItem", ldb.UpdateCalcItemHandler)
		calcRequest.GET("/startCalcItem/:id", ldb.StartCalculationItemHandler)
		calcRequest.GET("/stopCalcItem/:id", ldb.StopCalculationItemHandler)
		calcRequest.GET("/deleteCalcItem/:id", ldb.DeleteCalculationItemHandler)
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
	router.Static("/static", "./dist/static")  // load static files
	router.LoadHTMLGlob("./dist/*.html")  // render html template
	return router
}

func documentRouter() http.Handler {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return ""
	})) // customer console writing
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string);ok{
			utils.WriteError("", "", "", err)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{"Authorization", ""},
	}))
	router.Use(cors.Default())  // allow all cors
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/search_index.json", func(c *gin.Context) {
		c.File("./documents/_book/search_index.json")
	})
	router.GET("/images/db.png", func(c *gin.Context) {
		c.File("./documents/images/db.png")
	})
	router.GET("/GROUP.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "GROUP.html", nil)
	})
	router.GET("/ITEM.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ITEM.html", nil)
	})
	router.GET("/RESTFUL.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "RESTFUL.html", nil)
	})
	router.GET("/DATA.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "DATA.html", nil)
	})
	router.Static("/gitbook", "./documents/_book/gitbook")  // load static files
	router.LoadHTMLGlob("./documents/_book/*.html")  // render html template
	return router
}

func InitialDbServer(ip string, port int64, dbPath string, startReadConfigTime time.Time) error{
	checkResult, err := portInUse(port)
	if err != nil{
		return fmt.Errorf("%s: fail in checking port %d: %s", time.Now().Format(utils.TimeFormatString), port, err)
	}
	if checkResult != -1{
		// used
		return fmt.Errorf("%s: Failed to start web service: Port number %d is already occupied, process PID is %d, please consider using taskkill /f /pid %d to terminate the process", time.Now().Format("2006-01-02 15:04:05"), port, checkResult, checkResult)
	}
	checkResult1, err := portInUse(8087)
	if err != nil{
		return fmt.Errorf("%s: fail in checking port %d : %s", time.Now().Format(utils.TimeFormatString), port, err)
	}
	if checkResult1 != -1{
		// used
		return fmt.Errorf("%s: Failed to start web service: Port number %d is already occupied, process PID is %d, please consider using taskkill /f /pid %d to terminate the process", time.Now().Format("2006-01-02 15:04:05"), 8087, checkResult1, checkResult1)
	}
	gin.SetMode(gin.ReleaseMode)  // production
	address := ip + ":" + fmt.Sprintf("%d", port)  // base url of web server
	ldb := &db.LevelDb{}  // ldb pointer
	appServer := &http.Server{
		Addr:         address,
		Handler:      appRouter(dbPath, ldb),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	documentServer := &http.Server{
		Addr:         ip + ":8087",
		Handler:      documentRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	finalTime := time.Now()
	utils.WriteInfo(fmt.Sprintf("The system starts successfully, time consuming :%d ms", finalTime.Sub(startReadConfigTime).Milliseconds()))
	fmt.Printf("%s: launch web service successfully!: %s \n", time.Now().Format(utils.TimeFormatString), address)
	g.Go(func() error {
		err := appServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			utils.WriteError("", "", "", err.Error())  // write logs
		}
		return err
	})
	g.Go(func() error {
		err := documentServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			utils.WriteError("", "", "", err.Error())  // write logs
		}
		return err
	})
	g.Go(ldb.GetProcessInfo)  // monitor
	g.Go(ldb.Calc)  // calc goroutine
	if err := g.Wait(); err != nil {
		utils.WriteError("", "", "", err.Error())  // write logs
	}
	return nil
}
