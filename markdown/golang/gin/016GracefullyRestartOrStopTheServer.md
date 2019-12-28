# 优雅地重启或停止服务器

你想优雅地重启或停止 web 服务器吗？有一些方法可以做到这一点。我们可以使用 [fvbock/endless](https://github.com/fvbock/endless) 来替换默认的 `ListenAndServe`。更多详细信息，请参阅 issue [#296](https://github.com/gin-gonic/gin/issues/296)。

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 首页路由注册
	router.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "welcome to gin server")
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(endless.ListenAndServe(":8080", router))
}
```

## 替代方法

- [manners](https://github.com/braintree/manners)：可以优雅关机的 Go Http 服务器。
- [graceful](https://github.com/tylerb/graceful)：Graceful 是一个 Go 扩展包，可以优雅地关闭 http.Handler 服务器。
- [grace](https://github.com/facebookgo/grace)：Go 服务器平滑重启和零停机时间部署。

如果你使用的是 Go 1.8，可以不需要这些库！考虑使用 http.Server 内置的 [Shutdown()](https://golang.org/pkg/net/http/#Server.Shutdown) 方法优雅地关机. 请参阅 gin 完整的 [graceful-shutdown](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown) 示例。

## Golang版本1.8及以上处理方式

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 首页路由注册
	router.GET("/", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "welcome to gin server")
	})

	serve := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器(设置5秒的超时时间)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serve.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
```

## 目录

[BACK](../GolangGin.md)
