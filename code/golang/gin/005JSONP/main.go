package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 加载由glob模式标识的HTML文件并将结果与HTML渲染器关联起来
	router.LoadHTMLGlob("code/golang/gin/005JSONP/template/*")

	// 首页路由注册
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", nil)
	})

	// 获取用户邮箱信息路由注册
	router.GET("/email", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{
			"email": fmt.Sprintf("%d@foxmail.com", rand.Int()),
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
