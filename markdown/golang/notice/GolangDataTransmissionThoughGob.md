# Golang『用Gob传输数据』注意点

Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式；可以在 encoding 包中找到。这种格式的数据简称为 Gob （即 Go binary 的缩写）。类似于 Python 的 "pickle" 和 Java 的 "Serialization"。

Gob 通常用于远程方法调用（RPCs，参见 [15.9节](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.9.md) 的 rpc 包）参数和结果的传输，以及应用程序和机器之间的数据传输。 它和 JSON 或 XML 有什么不同呢？Gob 特定地用于纯 Go 的环境中，例如，两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。 Gob 不是可外部定义，语言无关的编码方式。因此它的首选格式是二进制，而不是像 JSON 和 XML 那样的文本格式。 Gob 并不是一种不同于 Go 的语言，而是在编码和解码过程中用到了 Go 的反射。

Gob 文件或流是完全自描述的：里面包含的所有类型都有一个对应的描述，并且总是可以用 Go 解码，而不需要了解文件的内容。

只有可被导出的字段会被编码，零值会被忽略。在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码。当源数据类型增加新字段后，Gob 解码客户端仍然可以以这种方式正常工作：解码客户端会继续识别以前存在的字段。并且还提供了很大的灵活性，比如在发送者看来，整数被编码成没有固定长度的可变长度，而忽略具体的 Go 类型。

假如在发送者这边有一个结构 T：

```go
type T struct { X, Y, Z int }
var t = T{X: 7, Y: 0, Z: 8}
```

而在接收者这边可以用一个结构体 U 类型的变量 u 来接收这个值：

```go
type U struct { X, Y *int8 }
var u U
```

在接收者中，X 的值是7，Y 的值是0（Y的值并没有从 t 中传递过来，因为它是零值）。

和 JSON 的使用方式一样，Gob 使用通用的 `io.Writer` 接口，通过 `NewEncoder` 函数创建 `Encoder` 对象并调用 `Encode()` 方法；相反的过程使用通用的 `io.Reader` 接口，通过 `NewDecoder()` 函数创建 `Decoder` 对象并调用 `Decode()` 方法。

下面将 `VCard` 结构体类型数据写进名为 `vcard.gob` 的文件作为例子。这会产生一个文本可读数据和二进制数据的混合，当你试着在文本编辑中打开的时候会看到。

```go
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// Using an gob encoder:
	file, err := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open the file error:%s\n", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encoding gob data error: %s\n", err)
		return
	}
	fmt.Printf("encoding successfully!")
}
```

运行程序会生成名称为 `vcard.gob` 的文件。

在下面示例中你会看到一个编解码，并且以字节缓冲模拟网络传输的简单例子：

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X    int
	Y    int
	Z    int
	Name string
}

type Q struct {
	X    *int32
	Y    *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.
	// 初始化编码器和解码器。
	// Normally enc and dec would be bound to network connections and the encoder and decoder would run in different processes.
	// 通常enc和dec是绑定到网络连接，编码器和解码器将在不同的进程中运行。

	// Stand-in for a network connection. (代替网络连接)
	var network bytes.Buffer

	// Will write to network. (将写入网络)
	enc := gob.NewEncoder(&network)

	// Will read from network. (将从网络读取)
	dec := gob.NewDecoder(&network)

	// Encode (send) the value. (编码(发送)值)
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatalf("encode error: %s\n", err)
	}

	// Decode (receive) the value. (解码(接收)值)
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatalf("decode error: %s\n", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
```

运行程序输出如下：

```shell
"Pythagoras": {3, 4}
```

## 练习：解码gob数据文件

写一个程序读取 vcard.gob 文件，解码并打印它的内容。

```go
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	// Open the file. (打开文件)
	file, err := os.Open("vcard.gob")
	if err != nil {
		log.Fatalf("open the file error: %s\n", err)
	}
	// Close the open file. (关闭打开的文件)
	defer file.Close()

	// Decode the gob content to structure. (将gob内容解码为结构)
	var vc VCard
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&vc)
	if err != nil {
		log.Fatalf("exercises the error: %s\n", err)
	}

	// Print the decoded structure data. (打印解码后的结构数据)
	fmt.Printf("exercises data structure: %+v\n", vc)
}
```

运行程序输出如下：

```shell
exercises data structure: {FirstName:Jan LastName:Kersschot Addresses:[0xc000064e10 0xc000064e40] Remark:none}
```

## 目录
[Back](../GolangNotice.md)
