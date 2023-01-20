package http_middleware

import (
	"gateway_server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func HTTPServiceHeaderTransform() gin.HandlerFunc {
	// support add/del/edit
	return func(c *gin.Context) {
		detail, err := util.GetServiceDetail(c)
		if err != nil {
			c.Abort()
			return
		}
		if detail.HttpRule.HeaderTransfor != "" {
			headerTransfor := strings.Split(detail.HttpRule.HeaderTransfor, "\n")
			for _, h := range headerTransfor {
				split := strings.Split(h, " ")
				if len(split) != 3 {
					continue
				}
				op := split[0]
				headerName := split[1]
				headerValue := split[2]
				switch op {
				case "add":
					c.Request.Header.Add(headerName, headerValue)
				case "del":
					c.Request.Header.Del(headerName)
				case "edit":
					c.Request.Header.Set(headerName, headerValue)
				}
			}
		}
		c.Next()
	}
}
