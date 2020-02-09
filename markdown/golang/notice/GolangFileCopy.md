# Golang『文件拷贝』注意点

如何拷贝一个文件到另一个文件？最简单的方式就是使用 `io` 包，请看下面的示例：

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("code/golang/read_writer_data/file_copy/target.txt", "code/golang/read_writer_data/file_copy/source.txt")
	fmt.Println("Copy done!")
}

// CopyFile is the a file copy function
func CopyFile(dstName string, srcName string) (written int64, err error) {
	// Open the file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// Close the open file
	defer src.Close()

	// Create a new file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	// Close the newly created file
	defer dst.Close()

	// Copy file
	return io.Copy(dst, src)
}
```
注意 `defer` 的使用：当打开dst文件时发生了错误，那么 `defer` 仍然能够确保 `src.Close()` 执行。如果不这么做，src文件会一直保持打开状态并占用资源。


## 目录
[Back](../GolangNotice.md)