package router

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/api"
)

func InitOAuthRouter(engine *gin.Engine) {
	group := engine.Group("/oauth")
	{
		group.POST("/tokens", api.Token)
	}
}
