package service

import "main.go/service/overWrite"

type ServiceGroup struct {
	OverWriteGroup overWrite.OverWriteGroup
}

var ServiceGroupApp = new(ServiceGroup)
