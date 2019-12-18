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
