package api

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"simple_gateway/model"
	"simple_gateway/util"
	"strings"
	"time"
)

type GatewayToken struct {
	Token     string `json:"token"`
	ExpireIn  int64  `json:"expire_in"`
	TokenType string `json:"token_type"`
}

func Token(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	split := strings.Split(auth, " ")
	if len(split) != 2 {
		util.RspError(c, util.CodeInvalidAuth, errors.New("认证格式出错"))
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		util.RspError(c, util.CodeInvalidAuth, errors.New("认证格式出错"))
		return
	}
	authStr := strings.Split(string(decodeString), ":")
	if len(authStr) != 2 {
		util.RspError(c, util.CodeInvalidAuth, errors.New("认证格式出错"))
		return
	}
	appId := authStr[0]
	secret := authStr[1]
	user := model.User{}
	all, _ := user.SelectAll()
	for _, v := range all {
		if v.AppID == appId && v.Secret == secret {
			token, err := util.GenerateToken(v.AppID)
			if err != nil {
				util.RspError(c, util.CodeServeBusy, errors.New("授权token出错"))
				return
			}
			gatewayToken := &GatewayToken{
				Token:     token,
				ExpireIn:  time.Now().Add(time.Hour * 24 * 7).Unix(),
				TokenType: "Bearer",
			}
			util.RspData(c, gatewayToken)
			return
		}
	}
}
