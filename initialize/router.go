package initialize

import (
	"github.com/gin-gonic/gin"
	"main.go/middleware"
	"main.go/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	//global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	//global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	OverWriterouter := router.RouterGroupApp.OverWrite
	OverWriteGroup := Router.Group("overwrite-api")
	PublicGroup := Router.Group("")

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		OverWriterouter.ScheduleOverWrite.InitMallCarouselIndexRouter(OverWriteGroup)
	}
	//global.GVA_LOG.Info("router register success")
	return Router
}
