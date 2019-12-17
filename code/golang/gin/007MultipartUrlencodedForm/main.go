package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册路由
	router.POST("/form-post", func(context *gin.Context) {
		message := context.PostForm("message")
		nickname := context.DefaultPostForm("nickname", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"status":   "posted",
			"message":  message,
			"nickname": nickname,
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
