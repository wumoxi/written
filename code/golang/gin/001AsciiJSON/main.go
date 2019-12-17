// 使用 AsciiJSON生成具有转义的非ASCII字符的 ASCII-onlyJSON
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{} {
			"lang": "Go 语言",
			"tag": "<br>",
		}
		// 输出：{"lang":"Go \u8bed\u8a00","tag":"\u003cbr\u003e"}
		context.AsciiJSON(http.StatusOK, data)
	})

	// 监听并在0.0.0.0:8080上启动服务
	log.Fatal(r.Run(":8080"))
}
