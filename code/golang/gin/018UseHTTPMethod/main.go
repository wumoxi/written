package main

import (
	"log"
	"written/code/golang/gin/018UseHTTPMethod/api/apiusers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 获取所有用户信息
	router.GET("/users", apiusers.GetUserAll())
	// 获取一个用户信息
	router.GET("/users/:id", apiusers.GetOne())
	// 添加一个用户
	router.POST("/users", apiusers.AddUser())
	// 修改一个用户完整信息
	router.PUT("/users/:id", apiusers.ChangeUser())
	// 修改一个用户部分信息
	router.PATCH("/users/:id", apiusers.ModifyUser())
	// 删除一个用户
	router.DELETE("/users/:id", apiusers.DeleteUser())
	// 获取用户资源的元数据
	router.HEAD("/users", apiusers.HeadUser())
	// 获取信息，关于用户资源的哪些属性是客户端可以改变的。
	router.OPTIONS("/users", apiusers.OptionsUser())

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
