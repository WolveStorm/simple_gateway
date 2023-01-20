package router

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/api"
)

func InitDashboardRouter(engine *gin.Engine) {
	group := engine.Group("/dashboard")
	{
		sa := api.DashboardApi{}
		group.GET("/panel_group_data", sa.PanelGroupData)
		group.GET("/flow_stat", sa.FlowStat)
		group.GET("/service_stat", sa.ServiceStat)

	}
}
