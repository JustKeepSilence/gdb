/*
creatTime: 2021/3/18
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

// middleware of gin for gdb web service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// authorization middleware
func (gdb *Gdb) authorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// for userLogin not need authorization
		if c.Request.URL.String() != "/page/userLogin" {
			if userName, token, ok := c.Request.BasicAuth(); !ok {
				c.AbortWithStatus(401)
			} else {
				userAgent := c.Request.Header.Get("User-Agent")
				if v, err := gdb.infoDb.Get([]byte(userName+"_token"+"_"+token+"_"+userAgent), nil); err != nil || v == nil {
					c.AbortWithStatus(401)
				} else {
					if token != fmt.Sprintf("%s", v) {
						c.AbortWithStatus(401)
					} else {
						c.Next()
					}
				}
			}
		} else {
			c.Next()
		}
	}
}

// writing logs and response body

func (gdb *Gdb) string(c *gin.Context, code int, formatter string, responseData []byte) {
	if level, ok := c.Request.Header["logLevel"]; ok {
		if level[0] == "0" {
			// info level
			if code == 200 {
				if c.Request.Method == "POST" {
					b, _ := ioutil.ReadAll(c.Request.Body)
					_ = gdb.writeLog(Info, c.Request.URL.String(), fmt.Sprintf("%s", b), "POST", "", c.Request.RemoteAddr)
				} else {
					_ = gdb.writeLog(Info, c.Request.URL.String(), c.Request.URL.String(), "GET", "", c.Request.RemoteAddr)
				}
			} else {
				if c.Request.Method == "POST" {
					b, _ := ioutil.ReadAll(c.Request.Body)
					_ = gdb.writeLog(Error, c.Request.URL.String(), fmt.Sprintf("%s", b), "POST", fmt.Sprintf("%s", responseData), c.Request.RemoteAddr)
				} else {
					_ = gdb.writeLog(Error, c.Request.URL.String(), c.Request.URL.String(), "GET", fmt.Sprintf("%s", responseData), c.Request.RemoteAddr)
				}
			}
		} else {
			// error
			if code != 200 {
				if c.Request.Method == "POST" {
					b, _ := ioutil.ReadAll(c.Request.Body)
					_ = gdb.writeLog(Error, c.Request.URL.String(), fmt.Sprintf("%s", b), "POST", fmt.Sprintf("%s", responseData), c.Request.RemoteAddr)
				} else {
					_ = gdb.writeLog(Error, c.Request.URL.String(), c.Request.URL.String(), "GET", fmt.Sprintf("%s", responseData), c.Request.RemoteAddr)
				}
			}
		}
	}
	if code == 500 {
		r, _ := Json.Marshal(ResponseData{
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
func (gdb *Gdb) setLogHeaderMiddleware(level logLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("logLevel", string(rune(level)))
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
