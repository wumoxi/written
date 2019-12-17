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
