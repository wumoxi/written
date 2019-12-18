package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用构建函数创建一个杜松子酒路由器(不包含任何中间件)
	router := gin.New()

	// 注册页面路由
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"info": "Welcome to Gin framework!",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
