# AsciiJSON

使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。它与JSON方法最大的区别在于对非Ascii字符不转义！

```go
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
```