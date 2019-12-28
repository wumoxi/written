# HTML 渲染

使用 LoadHTMLGlob() 或者 LoadHTMLFiles()

## 加载相同目录下所有HTML模板

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

	// 加载由glob模式标识的HTML文件并将结果与HTML渲染器关联起来
	router.LoadHTMLGlob("code/golang/gin/003HTMLReader/template/*")

	// 也可以使用LoadHTMLFiles逐个加载HTML模板文件
	// router.LoadHTMLFiles("template/template1.html", "template/template2.html")

	// 路由注册
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

其中 templates/index.tmpl 模板文件如下

```html
<html>
    <h1>
        {{ .title }}
    </h1>
</html>
```

运行WEB服务器并访问 `localhost:8080`，模板渲染如下图所示

```shell
$ go run server.go
```

![模板渲染](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_51.png)

## 使用不同目录下名称相同的模板

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

	// 加载由glob模式标识的HTML文件并将结果与HTML渲染器关联起来
	router.LoadHTMLGlob("code/golang/gin/003HTMLReader/template/**/*")

	// 帖子列表页
	router.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	// 用户列表页
	router.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

其中 templates/posts/index.tmpl 模板文件如下

```html
{{ define "posts/index.tmpl" }}
<html>
    <h1>{{ .title }}</h1>
    <p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

其中 templates/users/index.tmpl 模板文件如下

```html
{{ define "users/index.tmpl" }}
    <html>
    <h1>{{ .title }}</h1>
    <p>Using users/index.tmpl</p>
    </html>
{{ end }}
```

运行WEB服务器并访问 localhost:8080/posts/index 和 localhost:8080/users/index，模板渲染如下图所示

```shell
$ go run server.go
```
![localhost:8080\/posts\/index](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_52.png)
![localhost:8080\/users\/index](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_53.png)

## 自定义模板功能及自定义分隔符


```go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// formatAsDate模板时间解析函数
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 自定义模板分隔符
	router.Delims("{[{", "}]}")

	// 添加自定义模板解析函数
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// 加载一组HTML文件并将结果与HTML渲染器关联起来
	router.LoadHTMLFiles("code/golang/gin/003HTMLReader/template/raw.tmpl")

	// 注册首页路由
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Now(),
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

其中 raw.tmpl 模板如下所示

```html
Date: {[{ .now | formatAsDate }]}
```

运行WEB服务器并访问 localhost:8080，模板渲染如下图所示

```shell
$ go run server.go
```
![自定义模板功能及自定义分隔符](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_54.png)

## 目录

[BACK](../GolangGin.md)