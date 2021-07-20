// +build gdbServer

/*
creatTime: 2021/3/18
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.16
*/

package db

// middleware of gin for gdb web service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// authorization middleware
func (gdb *Gdb) authorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// for userLogin not need authorization
		if c.Request.URL.String() != "/page/userLogin" && !strings.Contains(c.Request.URL.String(), "/cmd/getCmdInfo") {
			if userName, token, ok := c.Request.BasicAuth(); !ok {
				c.AbortWithStatus(401)
			} else {
				if r, err := gdb.query("select token from user_cfg where userName='" + userName + "'"); err != nil || len(r) == 0 {
					c.AbortWithStatus(401)
				} else {
					if token != r[0]["token"] {
						c.AbortWithStatus(401)
					} else {
						// route permission check
						url := strings.Split(c.Request.URL.String(), "/")
						sub, obj, act := userName, strings.Title(url[len(url)-1]), c.Request.Method
						if ok, _ := gdb.e.Enforce(sub, obj, act); !ok {
							c.AbortWithStatus(401)
						} else {
							c.Request.Header.Add("userName", userName)
							c.Next()
						}
					}
				}
			}
		} else {
			c.Next()
		}
	}
}

// writing logs and response body

func (gdb *Gdb) string(c *gin.Context, code int, formatter string, responseData []byte, requestBody interface{}) {
	if level, ok := c.Request.Header["Loglevel"]; ok {
		b, _ := json.Marshal(requestBody)
		logMessage := logMessage{
			RequestUrl:    c.Request.URL.String(),
			RequestMethod: c.Request.Proto,
			UserAgent:     c.Request.UserAgent(),
			RequestBody:   fmt.Sprintf("%s", b),
			RemoteAddress: c.Request.RemoteAddr,
			Message:       "",
		}
		if level[0] == "Info" {
			// info level
			l := level[0]
			if code == 500 {
				l = "Error"
				logMessage.Message = strings.Replace(fmt.Sprintf("%s", responseData), "'", `"`, -1)
			}
			m, _ := json.Marshal(logMessage)
			_ = gdb.writeLog(l, fmt.Sprintf("%s", m), c.Request.Header.Get("userName"))
		} else {
			// error level
			if code != 200 {
				logMessage.Message = fmt.Sprintf("%s", responseData)
				m, _ := json.Marshal(logMessage)
				_ = gdb.writeLog("Error", fmt.Sprintf("%s", m), c.Request.Header.Get("userName"))
			}
		}
	}
	if code == 500 {
		r, _ := json.Marshal(ResponseData{
			Code:    500,
			Message: fmt.Sprintf("%s", responseData),
			Data:    "",
		})
		c.String(500, "%s", r)
	} else {
		c.String(code, formatter, responseData)
	}
}

// set logType to request headers
func (gdb *Gdb) setLogHeaderMiddleware(level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("logLevel", level)
		c.Next()
	}
}

func (gdb *Gdb) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (gdb *Gdb) writeLog(level, logMessage, requestUser string) error {
	if gdb.driverName == "mysql" {
		logMessage = strings.Replace(logMessage, `\`, `\\`, -1)
	}
	sqlString := "insert into log_cfg (logMessage, level, requestUser) values ('" + logMessage + "', '" + level + "','" + requestUser + "')"
	_, err := gdb.updateItem(sqlString)
	if err != nil {
		return err
	}
	return nil
}
