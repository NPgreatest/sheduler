package overWrite

import (
	"github.com/gin-gonic/gin"
	"main.go/model/common/response"
)

type overWriteApi struct {
}

func (o *overWriteApi) AddressList(c *gin.Context) {
	//token := c.GetHeader("token")
	//if err, userAddressList := mallUserAddressService.GetMyAddress(token); err != nil {
	//	global.GVA_LOG.Error("获取地址失败", zap.Error(err))
	//	response.FailWithMessage("获取地址失败:"+err.Error(), c)
	//} else if len(userAddressList) == 0 {
	//	response.OkWithData(nil, c)
	//} else {
	//	response.OkWithData(userAddressList, c)
	//}
	overWriteService.GetCarouselsForIndex()
	response.OkWithData([]int{1, 2, 3}, c)
}
