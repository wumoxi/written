# JSONP

```go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 加载由glob模式标识的HTML文件并将结果与HTML渲染器关联起来
	router.LoadHTMLGlob("code/golang/gin/005JSONP/template/*")

	// 首页路由注册
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", nil)
	})

	// 获取用户邮箱信息路由注册
	router.GET("/email", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{
			"email": fmt.Sprintf("%d@foxmail.com", rand.Int()),
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

其中index.tmpl内容如下，这个页面会生成一个内容为`SayHello`的`a`标签，点击这个`a`标签，会调用`email`接口，这个接口会返回一个随机的foxmail.com邮箱地址，如`1443635317331776148@foxmail.com`

```html
<html>
<a href="javascript: void(0);" onclick="SayHello('张三')">SayHello</a>
</html>

<script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
<script>
    function SayHello(name) {
        // Using YQL and JSONP
        $.ajax({
            url: "/email",

            // The name of the callback parameter, as specified by the YQL service
            jsonp: "callback",

            // Tell jQuery we're expecting JSONP
            dataType: "jsonp",

            // Tell YQL what we want and that we want JSON
            data: {
                format: "json"
            },

            // Work with the response
            success: function( response ) {
                console.log("hello, " + name + " your email address is " + response.email)
            }
        });
    }
</script>
```

运行WEB服务器并通过浏览器访问 `http://localhost:8080`，渲染页面如下图所示

```shell
$ go run server.go
```

![JSONP首页](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_58.png)

打开Chrome浏览器控制台，并点击`SayHello`标签，可以看到从WEB服务器接口`/email`获取到的数据！

![Chrome浏览器控制台](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_59.png)

再通过`network`标签也可以看到接口具体的响应数据

![接口具体的响应数据](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191217_61.png)

## 目录

[BACK](../gin-use.md)