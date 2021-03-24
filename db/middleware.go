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
func (gdb *Gdb) authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// for userLogin not need authorization
		if c.Request.URL.String() != "/page/userLogin" {
			if userName, token, ok := c.Request.BasicAuth(); !ok {
				c.AbortWithStatus(401)
			} else {
				if v, err := gdb.infoDb.Get([]byte(userName+"_token"), nil); err != nil || v == nil {
					c.AbortWithStatus(401)
				} else {
					if token != fmt.Sprintf("%s", v) {
						c.AbortWithStatus(401)
					}
				}
			}
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
	c.String(code, formatter, responseData)
}

// set logType to request headers
func (gdb *Gdb) setLogHeader(level logLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Add("logLevel", string(rune(level)))
	}
}
