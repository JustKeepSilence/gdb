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

// log middleware
func (gdb *Gdb) logsWriting() gin.HandlerFunc {
	return nil
}
