package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	// router.SecureJsonPrefix(")]}',\n")

	// 路由注册
	router.GET("/some-json", func(context *gin.Context) {
		// 语言切片数据定义
		languages := []string{"Golang", "javascript", "Java", "Python", "Lua", "C", "C++", "C#", "PHP"}

		// 将输出：while(1);["Golang","javascript","Java","Python","Lua","C","C++","C#","PHP"]
		context.SecureJSON(http.StatusOK, languages)
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
