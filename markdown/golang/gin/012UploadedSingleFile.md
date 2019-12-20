# 单文件上传

参考 issue [#774](https://github.com/gin-gonic/gin/issues/774) 和详细[示例代码](https://github.com/gin-gonic/examples/tree/master/upload-file/single).

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 定义上传目录
const uploadDir = "code/golang/gin/012UploadedSingleFile/public"

// 定义Host
const host = "localhost:8080"

func main() {
	// 用默认的中间件创建一个杜松子酒路由器(记录器logger和恢复recovery（无崩溃）中间件)
	router := gin.Default()

	// 为 multipart forms 设置较低的内存限制(默认是32Mib)
	router.MaxMultipartMemory = 8 << 20 // 8Mib

	// 设置静态资源访问目录
	router.Static("/", uploadDir)

	// 上传文件
	router.POST("/upload", func(context *gin.Context) {
		// 获取 POST 请求参数
		name := context.PostForm("name")
		email := context.PostForm("email")

		// 获取上传文件
		file, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"info": fmt.Sprintf("get form err: %s", err.Error()),
			})
			return
		}

		// 从上传临时文件夹移动到目录文件夹
		filename := filepath.Base(file.Filename)
		dst := fmt.Sprintf("%s/upload/%s", uploadDir, filename)
		visit := fmt.Sprintf("%s/upload/%s", host, filename)
		if err = context.SaveUploadedFile(file, dst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"info": fmt.Sprintf("upload file err: %s", err.Error()),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"info": fmt.Sprintf("File %s uploaded sucessfully with fields visit=%s name=%s and email=%s.", file.Filename, visit, name, email),
		})
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(router.Run(host))
}
```

运行WEB服务器并通过`curl`命令在命令行测试

```shell
$ go run server.go
```

使用 `curl` 命令通过命令行调用API接口 `/upload`

```shell
$ curl -XPOST "localhost:8080/upload" -F name="wumoxi" -F email="wu.shaohua@foxmail.com" -F file="@/usr/local/data/images/avatar.jpeg" -H "Content-Type: multipart/form-data"
{"info":"File avatar.jpeg upload sucessfully with fields visit=localhost:8080/upload/avatar.jpeg name=wumoxi and email=wu.shaohua@foxmail.com."}
```

可以看到响应数据中包含一个 `visit` 值，这个 `visit` 值就是可供上传后查看上传文件的URL地址，打开Chrome浏览器并访问该URL地址，可以预览上传后的图片文件。

![预览上传图像文件](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_65.png)

## 目录

[BACK](../GinUse.md)