package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"web_app/controllers"
	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs
	"web_app/logger"
	"web_app/middlewares"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.POST("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
		v1.GET("/posts2", controllers.GetPostListHandler2)
		//投票
		v1.POST("/vote", controllers.PostVoteController)
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
