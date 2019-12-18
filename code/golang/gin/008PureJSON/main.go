package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 提供 unicode 实体响应数据路由注册
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>hello, world!</b>",
		})
	})

	// 提供字面字符响应数据路由注册
	router.GET("/pure-json", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello, world!</b>",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
