package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"simple_gateway/global"
	"simple_gateway/initial"
	"simple_gateway/middleware"
	"simple_gateway/model"
	"simple_gateway/router"
	"syscall"
	"time"
)

func main() {
	initial.InitConfig(false)
	initial.InitAllZap()
	initial.InitMysql()
	initial.InitRedis()
	initial.ReplaceGinBinding("zh")
	model.SyncToRedis()
	defer initial.CloseConn()
	engine := gin.Default()
	store, err := redis.NewStore(10, "tcp", global.DebugFullConfig.RedisConfig.Host, global.DebugFullConfig.RedisConfig.Password, []byte("secret"))
	if err != nil {
		zap.S().Errorf("[Navi Gateway] err:%s", err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
	engine.Use(middleware.Cors(), middleware.AdminAuth())
	{
		router.InitAdminRouter(engine)
		router.InitServiceRouter(engine)
		router.InitDashboardRouter(engine)
		router.InitUserRouter(engine)
		router.InitOAuthRouter(engine)
	}
	server := http.Server{
		Addr:        global.DebugFullConfig.ServerConfig.Addr,
		Handler:     engine.Handler(),
		ReadTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		zap.S().Errorf("[Navi Gateway] dashboard server start error,ip:%s,err:%s", global.DebugFullConfig.ServerConfig.Addr, err.Error())
		return
	}
	zap.S().Infof("[Navi Gateway] dashboard server start addr:%s", global.DebugFullConfig.ServerConfig.Addr)
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign
}
