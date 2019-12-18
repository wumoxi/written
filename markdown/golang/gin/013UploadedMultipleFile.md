# 多文件上传

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
const uploadDir = "code/golang/gin/013UploadedMultipleFile/public"

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

		// 获取已解析的多部分表单，包括文件上传。
		form, err := context.MultipartForm()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"info": fmt.Sprintf("get form err:%s", err.Error()),
			})
			return
		}
		// 获取多文件上传表单域名上传文件
		files, ok := form.File["files[]"]
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{
				"info": "Unable to find upload form files!",
			})
			return
		}

		var visits []string

		// 从上传临时文件夹移动到目录文件夹
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			dst := fmt.Sprintf("%s/upload/%s", uploadDir, filename)
			visit := fmt.Sprintf("%s/upload/%s", host, filename)
			visits = append(visits, visit)
			if err = context.SaveUploadedFile(file, dst); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"info": fmt.Sprintf("upload file err: %s", err.Error()),
				})
				return
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"info":   "Multiple file uploaded successfully!",
			"visits": visits,
			"name":   name,
			"email":  email,
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
$ curl -XPOST "localhost:8080/upload" \
-F name="wumoxi" \
-F email="wu.shaohua@foxmail.com" \
-F "files[]=@/usr/local/data/images/gin-framework-logo-black.png" \
-F "files[]=@/usr/local/data/images/gin-framework-logo-blue.png" \
-H "Content-Type: multipart/form-data"
{"email":"wu.shaohua@foxmail.com","info":"Multiple file uploaded successfully!","name":"wumoxi","visits":["localhost:8080/upload/gin-framework-logo-black.png","localhost:8080/upload/gin-framework-logo-blue.png"]}
```

可以看到响应数据中包含一个 `visits` 值，这个 `visits` 值就是可供上传后查看上传文件的URL地址列表，打开Chrome浏览器并访问这些URL地址，可以预览上传后的图片文件。

![预览多文件上传](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_66.png)
![预览多文件上传](https://lucklit.oss-cn-beijing.aliyuncs.com/written/Snip20191218_67.png)

## 目录

[BACK](../gin-use.md)