package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟一些账户数据
var accounts = gin.H{
	"zhangsan": gin.H{"username": "zhangsan", "email": "zhangsan@foxmail.com", "phone": "13792892829"},
	"lisi":     gin.H{"username": "lisi", "email": "lisi@foxmail.com", "phone": "13792892830"},
	"wangwu":   gin.H{"username": "wangwu", "email": "wangwu@foxmail.com", "phone": "13792892831"},
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"zhangsan": "zhangsanpassword",
		"lisi":     "lisipassword",
		"wangwu":   "wangwupassword",
		"liliu":    "liliupassword",
	}))

	// /admin/secrets端点将触发 "localhost:8080/admin/secrets"
	authorized.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由BasicAuth中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)
		if account, ok := accounts[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "account": account})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "account": "No account -:("})
		}
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
