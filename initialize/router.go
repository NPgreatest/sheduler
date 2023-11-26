package initialize

import (
	"github.com/gin-gonic/gin"
	"main.go/middleware"
	"main.go/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors()) // 如需跨域可以打开
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
	return Router
}
