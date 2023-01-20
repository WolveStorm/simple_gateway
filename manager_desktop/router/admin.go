package router

import (
	"github.com/gin-gonic/gin"
	"simple_gateway/api"
)

func InitAdminRouter(engine *gin.Engine) {
	admin := engine.Group("/admin")
	{
		adminApi := api.AdminApi{}
		admin.POST("/login", adminApi.Login)
		admin.GET("/logout", adminApi.Logout)
		admin.POST("/change_pwd", adminApi.ChangePassword)
		admin.GET("/admin_info", adminApi.AdminInfo)
	}
}
