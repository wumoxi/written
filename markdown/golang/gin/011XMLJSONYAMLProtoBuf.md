# XML/JSON/YAML/ProtoBuf 渲染

```go
package main

import (
	"log"
	"net/http"
	"written/code/golang/gin/011XMLJSONYAMLProtoBuf/data"

	"github.com/gin-gonic/gin"
)

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	router.GET("/some-json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "This is gin framework for golang, it's great!",
			"status":  http.StatusOK,
		})
	})

	// 使用结构体为响应体JSON数据
	router.GET("/more-json", func(context *gin.Context) {
		// 你也可以使用一个结构体
		var message struct {
			Name    string `json:"name"`
			Message string
			Number  int
		}

		// 结构体字段赋值
		message.Name = "Tom"
		message.Message = "This is gin framework for golang, it's great!"
		message.Number = 2019
		// 注意 message.Name 在 JSON 中变成了 "name" 而不是 "Name" 这其决于message结构体Name字段的Tag标签声明
		// 将输出：{"name":"Tom","Message":"This is gin framework for golang, it's great!","Number":2019}
		context.JSON(http.StatusOK, message)
	})

	// 返回XML响应数据
	router.GET("/some-xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"message": "This is gin framework for golang, it's great!",
			"status":  http.StatusOK,
		})
	})

	// 返回YAML响应数据
	router.GET("/some-yaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "This is gin framework for golang, it's great!",
			"status":  http.StatusOK,
		})
	})

	// 返回Protobuf响应数据
	router.GET("/some-protobuf", func(context *gin.Context) {
		reqs := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf的具体定义写在 data/example.pf.go 文件中，这个文件是通过Protobuf生成的
		data := &example.Test{
			Label: &label,
			Reps:  reqs,
		}

		// 请注意，数据在响应中变为二进制数据
		// 将输出被 example.Test protobuf 序列化了的数据
		context.ProtoBuf(http.StatusOK, data)
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(":8080"))
}
```

其中生成 `data/example.pf.go` 文件的protobuf文件(example.proto)定义如下

```proto
syntax = "proto2";
package example;

message Test {
    required string label = 1;
    repeated int64 reps = 2;
}
```

进行到 `data` 目录，执行以下命令即可生成 `example.pf.go` 文件

```shell
$ protoc --go_out=plugins=grpc,import_path=example:. *.proto
```

运行WEB服务器并使用`curl`命令在命令行进行测试

```shell
$ go run server.go
```

使用 `curl` 命令调用API接口 `/some-json`

```shell
$ curl -XGET "localhost:8080/some-json"
{"message":"This is gin framework for golang, it's great!","status":200}
```

使用 `curl` 命令调用API接口 `/more-json`

```shell
$ curl -XGET "localhost:8080/more-json"
{"name":"Tom","Message":"This is gin framework for golang, it's great!","Number":2019}
```

使用 `curl` 命令调用API接口 `/some-xml`

```shell
$ curl -XGET "localhost:8080/some-xml" 
<map><message>This is gin framework for golang, it&#39;s great!</message><status>200</status></map>
```

使用 `curl` 命令调用API接口 `/some-yaml`

```shell
$ curl -XGET "localhost:8080/some-yaml"
message: This is gin framework for golang, it's great!
status: 200
```

使用 `curl` 命令调用API接口 `/some-protobuf`

```shell
$ curl -XGET "localhost:8080/some-protobuf"

test
```
## 目录

[BACK](../GinUse.md)