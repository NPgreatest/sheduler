package core

import (
	"fmt"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
