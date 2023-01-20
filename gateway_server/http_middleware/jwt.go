package http_middleware

import (
	"encoding/json"
	"gateway_server/cache"
	"gateway_server/global"
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		// 尝试从头部取出token
		auth := c.Request.Header.Get("Authorization")
		var match bool
		if auth != "" {
			split := strings.Split(auth, " ")
			if len(split) != 2 {
				util.RspError(c, util.CodeInvalidAuth, nil)
				c.Abort()
				return
			}
			token := split[1]
			if token != "" {
				users := cache.GetAllUSer()
				claim, err := util.VerifyToken(token)
				if err != nil {
					util.RspError(c, util.CodeInvalidAuth, nil)
					c.Abort()
					return
				}
				for _, u := range users {
					if u.AppID == claim.AppId {
						match = true
						user, _ := json.Marshal(u)
						c.Set(global.AppDetail, string(user))
						break
					}
				}
			}
		}
		// 如果网关开启了验证，且用户并没有输入token就需要进行拦截
		if detail.AccessControl.OpenAuth == 1 && !match {
			util.RspError(c, util.CodeInvalidAuth, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
