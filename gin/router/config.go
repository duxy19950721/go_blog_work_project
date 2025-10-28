package router

import (
	"blog/gin/biz"
	_ "blog/gin/docs"
	"blog/gin/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.PanicCatch())

	// 用户相关接口
	userGroup := router.Group("user")
	{
		userGroup.GET("/login", biz.Login)
		userGroup.GET("/register", biz.Registory)
	}

	// 文章相关接口
	postGroup := router.Group("post")
	{
		postGroup.POST("/create", biz.CreatePost)
		postGroup.GET("/query", biz.GetPostList)
		postGroup.POST("/update", biz.UpdatePost)
		postGroup.DELETE("/delete", biz.DeletePost)
	}

	// 评论相关接口
	commentGroup := router.Group("comment")
	{
		commentGroup.GET("/getListByPostId", biz.GetCommentListByPostId)
		commentGroup.POST("/create", biz.CreateComment)
	}

	registrySwagger(router)

	return router
}

func registrySwagger(router *gin.Engine) {
	// 添加docs路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置swagger.json路径
	router.StaticFile("swagger.json", "./docs/swagger.json")

	// 如果需要认证保护，可以这样配置：
	// authMiddleware := gin.BasicAuth(gin.Accounts{
	//     "admin": "1234",
	// })
	// router.GET("/docs/*any", authMiddleware, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
