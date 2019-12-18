# 使用BasicAuth中间件

```go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟一些账户数据
var accounts = gin.H{
	"zhangsan": gin.H{"username": "zhangsan", "email": "zhangsan@foxmail.com", "phone": "13792892829"},
	"lisi":     gin.H{"username": "lisi", "email": "lisi@foxmail.com", "phone": "13792892830"},
	"wangwu":   gin.H{"username": "wangwu", "email": "wangwu@foxmail.com", "phone": "13792892831"},
}

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		// 这里就当于用户名和密码的映射关系
		"zhangsan": "zhangsanpassword",
		"lisi":     "lisipassword",
		"wangwu":   "wangwupassword",
		"liliu":    "liliupassword",
	}))

	// /admin/secrets端点将触发 "localhost:8080/admin/secrets"
	authorized.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由BasicAuth中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)
		if account, ok := accounts[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "account": account})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "account": "No account -:("})
		}
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

运行WEB服务器并通过浏览器访问 `localhost:8080/admin/secrets`

```shell
$ go run server.go
```

![用户认证入口](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_70.png)

可以看到它提示我们输入用户名和密码，先试着输入一个在我们的账户列表中存在的用户如`zhangsan`，输入用户名为 `zhangsan`，密码为 `zhangsanpassword`，你会发现它通过了认证！

![使用BasicAuth中间件登录](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191219_72.png)

那直接再来，输入一个不存于我们的账户列表中的用户如`liliu`，你会发现它通过了认证！不过没有在我们的账户列表找到！

![使用BasicAuth中间件登录](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191219_73.png)


## 目录

[BACK](../gin-use.md)