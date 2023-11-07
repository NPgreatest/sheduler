package v1

import (
	"main.go/api/v1/overWrite"
)

type ApiGroup struct {
	OverWriteApiGroup overWrite.MallGroup
}

var ApiGroupApp = new(ApiGroup)
