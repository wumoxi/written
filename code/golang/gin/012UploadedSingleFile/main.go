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
