package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple_gateway/global/form"
	"simple_gateway/service"
	"simple_gateway/util"
)

type ServiceApi struct{}

// AddHTTPService 添加HTTP服务
func (s *ServiceApi) AddHTTPService(c *gin.Context) {
	var req form.AddHTTPServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.AddHTTPService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// UpdateHTTPService 更新HTTP服务
func (s *ServiceApi) UpdateHTTPService(c *gin.Context) {
	var req form.UpdateHTTPServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.UpdateHTTPService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// AddTCPService 添加TCP服务
func (s *ServiceApi) AddTCPService(c *gin.Context) {
	var req form.AddTCPServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddTCPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.AddTCPService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// UpdateTCPService 更新TCP服务
func (s *ServiceApi) UpdateTCPService(c *gin.Context) {
	var req form.UpdateTCPServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.UpdateTCPService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// AddGRPCService 添加GRPC服务
func (s *ServiceApi) AddGRPCService(c *gin.Context) {
	var req form.AddGRPCServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.AddGRPCService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// UpdateGRPCService 更新GRPC服务
func (s *ServiceApi) UpdateGRPCService(c *gin.Context) {
	var req form.UpdateGRPCServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.UpdateGRPCService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// ServiceDetail 服务明细
func (s *ServiceApi) ServiceDetail(c *gin.Context) {
	var req form.ServiceDetailReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] ServiceDetailReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	detail, err := service.ServiceDetail(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] get service detail error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, detail)
	return
}

// ServiceList 服务列表
func (s *ServiceApi) ServiceList(c *gin.Context) {
	var req form.ServiceListReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] ServiceListReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	list, err := service.ServiceList(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] get service detail error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, nil)
		return
	}
	util.RspData(c, list)
	return
}

// DeleteService 删除服务
func (s *ServiceApi) DeleteService(c *gin.Context) {
	var req form.DeleteHTTPServiceReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.DeleteService(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] add http error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

// ServiceStat 服务状态
func (s *ServiceApi) ServiceStat(c *gin.Context) {
	var req form.ServiceStatReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddHTTPServiceReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	stat, err := service.ServiceStat(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, stat)
}
