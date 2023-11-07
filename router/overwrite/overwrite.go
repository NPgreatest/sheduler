package overwrite

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type ScheduleOverWrite struct {
}

func (s *ScheduleOverWrite) InitMallCarouselIndexRouter(Router *gin.RouterGroup) {
	overWriteRouter := Router.Group("v1")
	var mallCarouselApi = v1.ApiGroupApp.OverWriteApiGroup
	{
		overWriteRouter.GET("get_schedule", mallCarouselApi.AddressList) // 获取首页数据
	}
}
