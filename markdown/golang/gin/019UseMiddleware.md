# 使用中间件(有明确留坑)

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将GIN_MODE 设置为 release.
	// gin.DefaultWriter 默认为标准输出
	router.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入500.
	router.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件
	router.GET("/benchmark", benchmarkLogger(), benchEndpoint)

	// 认证路由组
	// authorized := router.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样：
	authorized := router.Group("/")
	// 路由组中间件！在此例中，我们在 "authorized" 路由组中使用自定义的创建的 AuthRequired()中间件
	authorized.Use(AuthRequired())
	{
		// 登录
		authorized.POST("/login", loginEndpoint)
		// 注册
		authorized.POST("/submit", submitEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```