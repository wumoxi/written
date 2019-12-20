# Multipart/Urlencoded 绑定

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginForm用户登录认证类型
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册用户登录路由
	router.POST("/login", func(context *gin.Context) {
		// 你可以使用显式绑定声明方式进行绑定 multipart form:
		// context.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用ShouldBind方法进行自动绑定:
		var form = LoginForm{}
		// 在这种情况下，将自动选择合适的绑定
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"err": "form binding failed!",
			})
		} else {
			if form.Username == "wumoxi" && form.Password == "secretwumoxi" {
				// 登录成功
				context.JSON(http.StatusOK, gin.H{
					"status": "You are logged in",
				})
			} else {
				// 登录失败
				context.JSON(http.StatusUnauthorized, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过Postman进行API接口测试如下图所示

```shell
$ go run server.go
```

![PostmanAPI接口测试](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_63.png)

也可以通过 `curl` 命令进行命令行测试

```shell
$ curl --form username="wumoxi" --form password="secretwumoxi" -XPOST "localhost:8080/login"
{"status":"You are logged in"}
```

## 目录

[BACK](../GinUse.md)