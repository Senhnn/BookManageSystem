package main

import (
	"bookmanagesystem/config"
	"bookmanagesystem/controller"
	"bookmanagesystem/middleware/tls"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置文件初始化读取
	config.Init()

	httpsSvr := gin.New()
	httpSvr := gin.New()

	// 注册处理函数，注册中间件
	controller.HTTPHandlerInit(httpSvr)
	HTTPSHandlerInit(httpsSvr)

	// http服务器
	go func() {
		httpSvr.Run(config.ServerIp + ":" + config.ServerPort)
	}()

	// https服务器
	httpsSvr.RunTLS("localhost:20000", "key/server.crt", "key/server.key")
}

func HTTPSHandlerInit(httpsSvr *gin.Engine) {
	httpsSvr.Use(tls.TlsHandler())
}
