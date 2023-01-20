package api

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/service"
	"simple_gateway/util"
)

type DashboardApi struct{}

func (d *DashboardApi) PanelGroupData(c *gin.Context) {
	data, err := service.PanelGroupData()
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, data)
}

func (d *DashboardApi) ServiceStat(c *gin.Context) {
	data, err := service.DashboardServiceStat()
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, data)
}

func (d *DashboardApi) FlowStat(c *gin.Context) {
	stat, err := service.DashboardFlowStat()
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, stat)
}
