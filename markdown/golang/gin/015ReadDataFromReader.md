# 从reader读取数据

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
	router.GET("/some-data-from-reader", func(context *gin.Context) {
		// 发送一个GET类型的请求并获取响应数据
		resp, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || resp.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}

		// 获取响应reader，reps.Body其实是一个实现了ReadCloser接口的类型
		reader := resp.Body
		// 获取响应内容长度
		cl := resp.ContentLength
		// 获取响应数据类型
		ct := resp.Header.Get("Content-Type")
		// 附加请求头
		extraHeader := map[string]string{
			// 附加项: 附件并且文件名为gopher.png
			"Content-Disposition": "attachment; filename=gopher.png",
		}

		// 将指定的读取器(在这里就是reader)写入主体(body)流并更新HTTP代码。
		context.DataFromReader(http.StatusOK, cl, ct, reader, extraHeader)
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过浏览器访问 `localhost:8080/some-data-from-reader`

```shell
$ go run server.go
```

![从reader读取数据](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_69.gif)

可以看到当访问`localhost:8080/some-data-from-reader`这个URL的时候它会下载一张名为 `gopher.png` 的图片，这其中对应于服务端处理程序它所做的事情就是发起一个GET类型的HTTP请求，到 `https://raw.githubusercontent.com/gin-gonic/logo/master/color.png`这个地址，并且会获取到访问这个地址的响应数据，然后以附件的形式写入到`/some-data-from-reader` 这个API接口的响应体中，所示才会下载一个图片到本地！

## 目录

[BACK](../gin-use.md)

