# 只绑定URL查询字符串

`ShouldBindQuery` 和 `BindQuery` 函数只绑定`URL`查询参数而忽略`POST`数据。参阅[详细信息](https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017).

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 注册获取查询参数路由
	router.Any("/get-query-string", func(context *gin.Context) {
		var person Person
		if err := context.BindQuery(&person); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"binds": person,
			})
		}
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过Postman进行测试

```shell
$ go run server.go
```

通过Postman进行测试，使用GET方式调用API接口 `/get-query-string` 并添加查询参数如：`?name=张三&address=北京市东城区永外大街182号&mobile=13792928283`，
从下图结果可以看出，它仅仅绑定结构体中存在的字段，不存在的字段会被丢弃！

![Postman进行测试](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_91.png)

通过Postman进行测试，使用POST方式调用API接口 `/get-query-string`，从下图结果可以看出，POST方式的请求参数都会被丢弃，对的，这就是 `ShouldBindQuery` 和 `BindQuery` 方法的功能本质！

![Postman进行API接口测试](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191220_92.png)

## 目录

[BACK](../GolangGin.md)