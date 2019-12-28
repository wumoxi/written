# 控制日志输出颜色

默认情况下，控制台上输出的日志应根据检测到的TTY进行着色。

## 禁用日志的颜色

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 禁用日志的颜色
	gin.DisableConsoleColor()

	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册测试URI
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行HTTP服务端程序，会看到如下信息

```text
$ go run server.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

当有请求被处理时，例如访问注册路由 `localhost:8080/ping` logger会默认在标准输出记录请求处理日志

![禁用日志的颜色](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_5.png)

## 始终为日志着色

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 强制记录日志的颜色
	gin.ForceConsoleColor()

	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册测试URI
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行HTTP服务端程序，会看到如下信息

```text
$ go run server.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

当有请求被处理时，例如访问注册路由 `localhost:8080/ping` logger会默认在标准输出记录请求处理日志

![强制记录日志的颜色](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_4.png)

## 目录

[BACK](../GolangGin.md)