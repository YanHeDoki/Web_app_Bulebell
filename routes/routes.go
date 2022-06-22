package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)

	}

	//r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//
	//	//判断用户是否登录用请求头中是否有有效的jwt
	//	c.Request.Header.Get("Authorization")
	//	isLoign := true
	//	if isLoign {
	//		//如果登录用户
	//		c.String(http.StatusOK, "pong")
	//	} else {
	//		//否则返回请登录
	//		c.String(http.StatusOK, "请登录")
	//	}
	//})

	return r
}
