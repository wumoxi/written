# 不使用默认的中间件

使用构造函数 `gin.New()` 用于替代 `gin.Default()`  函数，创建一个纯净的 `Gin` 引擎，它不包含任何默认的中间件！

```go
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
```

运行WEB服务器并通过浏览器访问 `localhost:8080` 可以看到一个首页JSON响应体预览

```shell
$ go run server.go
```

![不使用默认的中间件](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_68.png)

## 目录

[BACK](../GolangGin.md)