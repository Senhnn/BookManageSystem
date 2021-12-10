package main

import (
	"bookmanagesystem/controller"
	"bookmanagesystem/middleware/tls"
	"github.com/gin-gonic/gin"
)

func main() {
	httpsSvr := gin.New()
	httpSvr := gin.New()

	// 注册处理函数，注册中间件
	HTTPHandlerInit(httpSvr)
	HTTPSHandlerInit(httpsSvr)

	// http服务器
	go func() {
		httpSvr.Run("localhost:10000")
	}()

	// https服务器
	httpsSvr.RunTLS("localhost:20000", "key/server.crt", "key/server.key")
}

func HTTPHandlerInit(httpSvr *gin.Engine) {
	httpSvr.POST("/adduser", controller.AddUser)
}

func HTTPSHandlerInit(httpsSvr *gin.Engine) {
	httpsSvr.Use(tls.TlsHandler())
}
