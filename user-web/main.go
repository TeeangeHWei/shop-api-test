package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop-api/user-web/initialize"
)

func main() {
	// 1. 初始化router
	Router := initialize.Routers()

	zap.S().Info("启动服务器，端口:%d", 8024)

	port := 8021
	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
