package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"simple_gateway/global"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(global.AdminInfoSessionKey) == "" {
			c.Abort()
		}
		c.Next()
	}
}
