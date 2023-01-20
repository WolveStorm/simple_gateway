package api

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"simple_gateway/dto"
	"simple_gateway/global"
	"simple_gateway/global/form"
	"simple_gateway/service"
	"simple_gateway/util"
	"time"
)

type AdminApi struct{}
type AdminSession struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}
type AdminOutPut struct {
	Id           int       `json:"id"`
	Username     string    `json:"username"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

func (a *AdminApi) Login(c *gin.Context) {
	var req form.AdminLoginInput
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AdminLoginInput Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	info, err := service.Login(req.Username, req.Password)
	if err != nil {
		util.RspError(c, util.CodeInvalidPassword, nil)
		return
	}
	s := AdminSession{
		Id:        int(info.ID),
		Username:  info.UserName,
		LoginTime: time.Now(),
	}
	jsonS, err := json.Marshal(s)
	if err != nil {
		util.RspError(c, util.CodeServeBusy, nil)
		return
	}
	session := sessions.Default(c)
	session.Set(global.AdminInfoSessionKey, string(jsonS))
	session.Save()
	util.RspData(c, dto.AdminLoginOutput{Token: global.AdminInfoSessionKey})
}

func (a *AdminApi) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(global.AdminInfoSessionKey)
	session.Save()
	util.RspData(c, nil)
}

func (a *AdminApi) ChangePassword(c *gin.Context) {
	var req form.ChangePasswordReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] ChangePasswordReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	err := service.ChangePassword(req.Password)
	if err != nil {
		zap.S().Warnf("[Navi Gateway] change pwd error,err:%s", err.Error())
		util.RspError(c, util.CodeServeBusy, nil)
		return
	}
	util.RspData(c, nil)
}

func (a *AdminApi) AdminInfo(c *gin.Context) {
	var req form.AdminInfoReq
	if err := c.ShouldBind(&req); err != nil {
		zap.S().Warnf("[Navi Gateway] AdminInfoReq Bind Error,err:%s", err.Error())
		util.RspBindingError(c, err)
		return
	}
	session := sessions.Default(c)
	infoStr := session.Get(req.Token)
	info := AdminSession{}
	if str, ok := infoStr.(string); ok {
		err := json.Unmarshal([]byte(str), &info)
		if err != nil {
			util.RspError(c, util.CodeServeBusy, nil)
			return
		}
		out := AdminOutPut{
			Id:           info.Id,
			Username:     info.Username,
			LoginTime:    info.LoginTime,
			Avatar:       "https://s1.ax1x.com/2022/04/22/LgIprR.jpg",
			Introduction: "Natus Vincere Super Admin",
			Roles:        []string{"admin"},
		}
		util.RspData(c, out)
	} else {
		util.RspData(c, nil)
	}
}
