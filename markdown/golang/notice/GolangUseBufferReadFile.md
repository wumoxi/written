# Golang『用 buffer读取文件』注意点

在下面的例子中，我们结合使用了缓冲读取文件和命令行 flag 解析这两项技术。如果不加参数，那么你输入什么屏幕就打印什么。参数被认为是文件名，如果文件存在的话就打印文件内容到屏幕。在命令行执行 `go run cat.go README.md` 测试输出。 

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buff, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buff)
	}
	return
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: error: %s\n", os.Args[0], flag.Arg(i), err)
			return
		}
		cat(bufio.NewReader(file))
	}
}
```

在 [12.6](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.6.md) 章节，我们将看到如何使用缓冲写入。

## 练习：为显示文件内容添加行号

扩展上面的例子，使用 `flag` 添加一个选项，目的是为每一行头部加入一个行号。使用 `go run cat.go -n README.md` 测试输出。

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var isShowLineNumber = flag.Bool("n", false, "show file content of line number")

func cat(r *bufio.Reader) {
	var number int64
	for {
		buff, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *isShowLineNumber {
			number++
			fmt.Fprintf(os.Stdout, "%5d %s", number, buff)
		} else {
			fmt.Fprintf(os.Stdout, "%s", buff)
		}
	}
	return
}

func main() {
	flag.PrintDefaults()
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: error: %s\n", os.Args[0], flag.Arg(i), err)
			return
		}
		cat(bufio.NewReader(file))
	}
}
```

