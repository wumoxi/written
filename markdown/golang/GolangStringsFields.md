# Golang切割包含不定个数空格的字符串怎么处理？

遇到这么一个情况，比如说一个字符串内容为 `中国人    美国人  英国人 法国人`，要把它切割为一个数组，最常用的字符串切割方法应该是 `strings.Split`，可是这个字符串中的子字符串与子字符串之间的空格数量不一致。要怎么处理？下面先来看看使用 `strings.Split` 处理会是个什么结果。

## 错误的处理方式

如果使用`strings.Split`方法，你将会得到这样一个数组切片

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "中国人    美国人  英国人 法国人"
	sli := strings.Split(s, " ")
	fmt.Printf("slice: %+v, len: %d\n", sli, len(sli))
}
```

输出结果如下，这其实不是我想要的，看看这长度明明预想的是4，结果却出来一个8，原因是这个数组切片它的item包含空格元素值! 这么来说要对这个数组切片进行去重，这不好吧，还要这么麻烦，有木有更好的办法？哦，对了，strings 包好像有一个 `Fields` 方法，对的，它可以做到！

```shell
slice: [中国人    美国人  英国人 法国人], len: 8
```

## 正确的处理方式

```shell
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "中国人    美国人  英国人 法国人"
	sli := strings.Fields(s)
	fmt.Printf("slice: %+v, len: %d\n", sli, len(sli))
}
```

输出结果如下，这很好！不错，就应该是这个样子的！

```shell
slice: [中国人 美国人 英国人 法国人], len: 4
```

## 关于`strings.Fields`方法的说明

`strings.Fields(s)` 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。
