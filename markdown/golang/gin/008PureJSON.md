# PureJSON

通常，JSON 使用 unicode 编码替换特殊 HTML 字符，例如 < 变为 \ u003c。如果要按字面对这些字符进行编码，则可以使用 PureJSON。Go 1.6 及更低版本无法使用此功能。

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

	// 提供 unicode 实体响应数据路由注册
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>hello, world!</b>",
		})
	})

	// 提供字面字符响应数据路由注册
	router.GET("/pure-json", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello, world!</b>",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并使用`curl`命令在命令行进行测试

```shell
$ go run server.go
```

使用 `curl` 命令调用 `/json` API接口

```shell
$ curl -XGET "localhost:8080/json"          
{"html":"\u003cb\u003ehello, world!\u003c/b\u003e"}
```

使用 `curl` 命令调用 `/pure-json` API接口

```shell
$ curl -XGET "localhost:8080/pure-json"   
{"html":"<b>hello, world!</b>"}
```

通过上面的API接口调用可以看到，使用`JSON`方法时，会使用 `unicode` 对实体标记符进行编码，而使用`PureJSON`方法时，会按实体标记字面字符进行原样编码(也就是原样返回)！从方法命名也可以看出它的具体功能，`Pure` 单词的本意为 `纯`， 那 `PureJSON` 也就意味着 `纯JSON` 不进行任何后端处理！

## 目录

[BACK](../gin-use.md)
