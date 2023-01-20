package router

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/api"
)

func InitServiceRouter(engine *gin.Engine) {
	group := engine.Group("/service")
	{
		sa := api.ServiceApi{}
		group.POST("service_add_http", sa.AddHTTPService)
		group.POST("service_update_http", sa.UpdateHTTPService)
		group.POST("service_add_tcp", sa.AddTCPService)
		group.POST("service_update_tcp", sa.UpdateTCPService)
		group.POST("service_add_grpc", sa.AddGRPCService)
		group.POST("service_update_grpc", sa.UpdateGRPCService)
		group.GET("service_detail", sa.ServiceDetail)
		group.GET("/service_list", sa.ServiceList)
		group.GET("/service_delete", sa.DeleteService)
		group.GET("/service_stat", sa.ServiceStat)
	}
}
