# Query和PostForm

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 路由注册
	router.POST("/post", func(context *gin.Context) {
		// 获取请求参数
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")

		// 以JSON方式返回获取到的请求参数响应数据
		context.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过`curl`命令进行命令行测试

```shell
$ go run server.go
```

通过`curl`命令在命令行调用API接口`/post`

```shell
$ curl --form name="ginner" --form message="This is framework for golang" -XPOST "localhost:8080/post?id=8859&page=1"
{"id":"8859","message":"This is framework for golang","name":"ginner","page":"1"} 
```

从这个Demo中可以看到通过`Query`方法和`PostForm`方法，可以进行`GET`请求参数的获取和`POST`请求参数的获取，当然`Query`也可以只用于`HTTP-GET`请求类型的参数获取，`PostForm`也可以只用于`HTTP-POST`请求类型的参数获取，不过这个Demo演示了，既存在GET请求参数也存在POST请求参数的获取，这就是通过`Query`和`PostForm`方法结合得以实现！

## 目录

[BACK](../gin-use.md)
