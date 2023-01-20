package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple_gateway/global/form"
	"simple_gateway/service"
	"simple_gateway/util"
)

type UserApi struct{}

func (a UserApi) AppList(c *gin.Context) {
	var req form.AppListReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AppListReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	list, err := service.UserList(req)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] get service detail error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, nil)
		return
	}
	util.RspData(c, list)
	return
}

func (a UserApi) AppDetail(c *gin.Context) {
	var req form.UserDetailReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] UserDetailReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	detail, err := service.AppDetail(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, detail)
}

func (a UserApi) AppStat(c *gin.Context) {
	var req form.UserStatReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] UserStatReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	stat, err := service.UserStat(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, stat)
}

func (a UserApi) AppDelete(c *gin.Context) {
	var req form.DeleteUserReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] DeleteUserReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.DeleteUser(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

func (a UserApi) AppAdd(c *gin.Context) {
	var req form.AddUserReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AddUserReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.AddUser(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}

func (a UserApi) AppUpdate(c *gin.Context) {
	var req form.UpdateUserReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] UpdateUserReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.UpdateUser(req)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, err)
		return
	}
	util.RspData(c, nil)
}
