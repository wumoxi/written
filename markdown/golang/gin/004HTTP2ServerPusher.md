# HTTP2 server 推送

http.Pusher 仅支持 go1.8+。 更多信息，请查阅 [golang blog](https://blog.golang.org/h2push)。

```go
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定义模板渲染器
var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	router.Static("/assets", "code/golang/gin/004HTTP2ServerPusher/assets")

	// 使用自定义模板渲染器
	router.SetHTMLTemplate(html)

	router.GET("/", func(context *gin.Context) {
		if pusher := context.Writer.Pusher(); pusher != nil {
			// 使用pusher.Push()方法做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %s\n", err)
			}
		}
		context.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并在浏览器访问 `localhost:8080`，渲染页面如下所示


```shell
$ go run server.go
```

![HTTP2 server 推送](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_57.png)