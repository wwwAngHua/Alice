package main

import (
	"Alice/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRoute(route *gin.Engine) *gin.Engine {
	// 支持多Robot分组
	robot := route.Group("/robot")
	{
		// Alice
		alice := robot.Group("/alice")
		{
			alice.POST("/chat", controller.Chat)     //  交谈
			alice.POST("/upload", controller.Upload) //  上传文件
		}
		// 学习助手
		lean := robot.Group("/lean")
		{
			lean.POST("/chat", controller.LeanChat) // 交谈
			lean.POST("/upload", controller.Upload) // 上传文件
		}
		// 帮助
		help := robot.Group("/help")
		{
			help.POST("/chat", controller.HelpChat) // 交谈
			help.POST("/upload", controller.Upload) // 上传文件
		}
		// 免费图床
		image := robot.Group("/image")
		{
			image.POST("/chat", controller.ImageChat)     // 交谈
			image.POST("/upload", controller.ImageUpload) // 上传文件
		}
		// 电影解析
		movie := robot.Group("/movie")
		{
			movie.POST("/chat", controller.MovieChat) // 交谈
			movie.POST("/upload", controller.Upload)  // 上传文件
		}
	}

	// No router
	route.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 Not Found")
	})

	return route
}
