# Golang『XML数据格式』注意点

下面是与 [12.9](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.9.md) 节 JSON 例子等价的 XML 版本：

```xml
<Person>
    <FirstName>沫汐</FirstName>
    <LastName>武</LastName>
</Person>
```

如同 json 包一样，也有 `Marshal()` 和 `UnMarshal()` 从 XML 中编码和解码数据；但这个更通用，可以从文件中读取和写入（或者任何实现了 io.Reader 和 io.Writer 接口的类型）

和 JSON 的方式一样，可以将结构数据序列化为 XML 数据，或者将 XML 数据反序列化为结构数据；这些可以在例子 [15.8(twitter_status.go)](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_15/twitter_status.go) 中看到。

encoding/xml 包实现了一个简单的 XML 解析器（SAX），用来解析 XML 数据内容。下面的例子说明如何使用解析器：

```go
package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	input := `<Person><FirstName>沫汐</FirstName><LastName>武</LastName></Person>`
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)

	for t, err := p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)

			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			fmt.Printf("%v is can't process", token)
		}
	}
}
```

运行程序输出如下：

```shell
Token name: Person
Token name: FirstName
This is the content: 沫汐
End of token
Token name: LastName
This is the content: 武
End of token
End of token
```

包中定义了若干 XML 标签类型：StartElement，Chardata（这是从开始标签到结束标签之间的实际文本），EndElement，Comment，Directive 或 ProcInst。


包中同样定义了一个结构解析器：`NewParser` 方法持有一个 io.Reader（这里具体类型是 strings.NewReader）并生成一个解析器类型的对象。还有一个 `Token()` 方法返回输入流里的下一个 XML token。在输入流的结尾处，会返回（nil，io.EOF）

XML 文本被循环处理直到 `Token()` 返回一个错误，因为已经到达文件尾部，再没有内容可供处理了。通过一个 type-switch 可以根据一些 XML 标签进一步处理。Chardata 中的内容只是一个 []byte，通过字符串转换让其变得可读性强一些。

## 目录
[Back](../GolangNotice.md)