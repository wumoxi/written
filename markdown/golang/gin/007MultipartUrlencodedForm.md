# Multipart/Urlencoded 表单

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

	// 注册路由
	router.POST("/form-post", func(context *gin.Context) {
		message := context.PostForm("message")
		nickname := context.DefaultPostForm("nickname", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"status":   "posted",
			"message":  message,
			"nickname": nickname,
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过Postman进行API接口测试如下图所示

```shell
$ go run server.go
```

![PostmanAPI接口测试](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_64.png)

也可以通过 `curl` 命令进行命令行测试

```shell
$ curl --form  message="My name is WuMoxi" --form nickname="wumoxi" -XPOST "localhost:8080/form-post"
{"message":"My name is WuMoxi","nickname":"wumoxi","status":"posted"}
```

## 目录

[BACK](../gin-use.md)