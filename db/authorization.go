/*
creatTime: 2021/3/15
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
goVersion: 1.15.3
*/

package db

import (
	"github.com/gin-gonic/gin"
)

// authorization middleware, for details see:
func (gdb *Gdb) authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// for userLogin not need authorization
		if c.Request.URL.String() != "/page/userLogin" {
			if userName, passWord, ok := c.Request.BasicAuth(); !ok {
				c.AbortWithStatus(401)
			} else {
				if v, err := gdb.infoDb.Get([]byte(userName), nil); err != nil || v == nil {
					c.AbortWithStatus(401)
				} else {
					userInfo := userInfo{}
					_ = Json.Unmarshal(v, &userInfo)
					if userInfo.PassWord != passWord {
						c.AbortWithStatus(401)
					}
				}
			}
		}
	}
}
