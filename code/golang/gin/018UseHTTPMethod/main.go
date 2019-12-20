package main

import (
	"log"
	"written/code/golang/gin/018UseHTTPMethod/router/users"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册用户管理模块API路由
	routerusers.Users(router)

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
