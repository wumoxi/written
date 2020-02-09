# Golang『用切片读写文件』注意点

切片提供了Go中处理I/O缓冲的标准方式，下面代码为 `cat` 函数的第二个版本，在一个切片缓冲内使用无限 for 循环（直到文件尾部 EOF）读取文件，并写入到标准输出（`os.Stdout`）。

```go
func cat(f *os.File) {
	const NBUF = 512
	var buff [NBUF]byte
	for {
		switch nr, err := f.Read(buff[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:   // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buff[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
```

上面的代码是下面要展示程序的一部分，使用了 os 包中的 `os.File` 和 `Read` 方法；请看下面的完整示例：

```go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buff [NBUF]byte
	for {
		switch nr, err := f.Read(buff[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:  // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buff[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if file == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error: %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(file)
		file.Close()
	}
}
``` 

这个示例和[Golang『用 buffer读取文件』注意点](https://github.com/wumoxi/written/blob/master/markdown/golang/notice/GolangUseBufferReadFile.md#golang%E7%94%A8-buffer%E8%AF%BB%E5%8F%96%E6%96%87%E4%BB%B6%E6%B3%A8%E6%84%8F%E7%82%B9) 具有同样的功能。


## 目录
[Back](../GolangNotice.md)