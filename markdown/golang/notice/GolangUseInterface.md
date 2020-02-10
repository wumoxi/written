# Golang『使用接口的实际例子：fmt.Fprintf』注意点

下面示例很好的阐述了 io 包中的接口概念。请看下面示例程序：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world - unbuffered")

	// buffered: os.Stdout implements io.Writer
	buff := bufio.NewWriter(os.Stdout)

	// and now so does buff.
	fmt.Fprintf(buff, "%s\n", "hello world - buffered")

	// Flush buffer data to buff
	buff.Flush()
}
```

运行程序输出如下：

```shell
hello world - unbuffered
hello world - buffered
```

下面是 `fmt.Fprintf()` 函数的实际签名

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

其不是写入一个文件，而是写入一个 `io.Writer` 接口类型的变量，下面是 `Writer` 接口在 io 包中的定义：

```go
// Writer is the interface that wraps the basic Write method.
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

`fmt.Fprintf()` 依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了 `io.Writer` 接口。`Fprintf()` 能够写入任何类型，只要其实现了 `Write` 方法，包括 `os.Stdout`、文件（例如os.File）、网络连接、通道等等，同样的也可以使用 bufio 包中缓冲写入。bufio 包中定义了 `type Writer struct{...}`。

bufio.Writer 实现了 Write 方法：

```go
// Write writes the contents of p into the buffer.
// It returns the number of bytes written.
// If nn < len(p), it also returns an error explaining
// why the write is short.
func (b *Writer) Write(p []byte) (nn int, err error) {
	for len(p) > b.Available() && b.err == nil {
		var n int
		if b.Buffered() == 0 {
			// Large write, empty buffer.
			// Write directly from p to avoid copy.
			n, b.err = b.wr.Write(p)
		} else {
			n = copy(b.buf[b.n:], p)
			b.n += n
			b.Flush()
		}
		nn += n
		p = p[n:]
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}
```

它还有一个工厂函数：传给它一个 `io.Writer` 接口类型的参数，它会返回一个带缓冲的 `bufio.Writer` 类型的 `io.Writer` 接口类型：

```go
// NewWriter returns a new Writer whose buffer has the default size.
func NewWriter(w io.Writer) *Writer {
	return NewWriterSize(w, defaultBufSize)
}
```

其适合任何形式的缓冲写入。

在缓冲写入的最后千万不要忘了使用 `Flush()`，否则最后的输出不会被写入。

在 15.2~15.8章节，我们将使用 `fmt.Fprint` 函数向 `http.ResponseWriter` 写入，其同样实现了 `io.Writer` 接口。

## 练习：截取行部分数据写入到另一个文件

下面的代码有一个输入文件 `goprogram`，然后以每一行为单位读取，从读取的当前行中截取第 3 到第 5 的字节写入另一个文件。然而当你运行这个程序，输出的文件却是个空文件。找出程序逻辑中的 bug，修正它并测试。

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("code/golang/read_writer_data/use_interface/exercises/12.7/goprogram.txt")
	outputFile, _ := os.OpenFile("code/golang/read_writer_data/use_interface/exercises/12.7/goprogramT.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		fmt.Println(inputString)
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
```

这个程序的 bug 在于没有将写入到缓冲区的数据刷入到文件，所以它会产生一个空的文件，修复这个 bug 并不难，主要在于我们对 `bufio.Writer` 类型的理解和熟知，在使用 `buffer.Writer` 的时候，一定要记得将写入到缓冲区内的数据，使用 `Flush()` 方法，刷入到文件(*os.File类型)。所以添加一行代码即可解决掉这个 bug:

```go
// Flush the buffer data to a file
outputWriter.Flush()
```

修改后的完整文件如下所示：

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("code/golang/read_writer_data/use_interface/exercises/12.7/goprogram.txt")
	outputFile, _ := os.OpenFile("code/golang/read_writer_data/use_interface/exercises/12.7/goprogramT.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		fmt.Println(inputString)
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Flush the buffer data to a file
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")
}
```

## 目录
[Back](../GolangNotice.md)
