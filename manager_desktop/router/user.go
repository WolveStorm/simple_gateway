package router

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/api"
)

func InitUserRouter(engine *gin.Engine) {
	group := engine.Group("/app")
	{
		ua := &api.UserApi{}
		group.GET("app_list", ua.AppList)
		group.GET("app_detail", ua.AppDetail)
		group.GET("app_stat", ua.AppStat)
		group.GET("app_delete", ua.AppDelete)
		group.POST("app_add", ua.AppAdd)
		group.POST("app_update", ua.AppUpdate)

	}
}
