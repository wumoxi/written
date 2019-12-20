package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册获取查询参数路由
	router.Any("/get-query-string", func(context *gin.Context) {
		var person Person
		if err := context.BindQuery(&person); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"binds": person,
			})
		}
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
