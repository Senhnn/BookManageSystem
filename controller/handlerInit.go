package controller

import (
	"bookmanagesystem/middleware"
	"github.com/gin-gonic/gin"
)

func HTTPHandlerInit(httpSvr *gin.Engine) {
	// 使用recover中间件
	httpSvr.Use(gin.Recovery())

	// 需要鉴权的路由组
	secureGroup := httpSvr.Group("/")
	{
		// 使用jwt中间件
		secureGroup.Use(middleware.JwtTokenAuth())
		secureGroup.POST("/adduser", UserAdd)
	}

	// 无须鉴权的路由组
	simpleGroup := httpSvr.Group("/")
	{
		simpleGroup.POST("register", UserRegister)
	}
}
