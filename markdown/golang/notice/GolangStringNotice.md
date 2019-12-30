# Golang『字符串』注意点

## 反转字符串

### 通过字节切片反转字符串

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "中国人民解放军8341"
	fmt.Printf("reverse before of string: %s\n", s)
	fmt.Printf("reverse after of string: %s\n", ReverseStr([]byte(s)))
}

// ReverseStr反转字符串
func ReverseStr(p []byte) string {
	s := make([]rune, len(p))
	for len(p) > 0 {
		r, size := utf8.DecodeLastRune(p)
		s = append(s, r)
		p = p[:len(p)-size]
	}
	return string(s)
}
```

程序输出如下

```shell
reverse before of string: 中国人民解放军8341
reverse after of string: 1438军放解民人国中
```

### 通过字符串反转字符串

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "中国人民解放军8341"
	fmt.Printf("reverse before of string: %s\n", s)
	fmt.Printf("reverse after of string: %s\n", ReverseStr(s))
}

// ReverseStr反转字符串
func ReverseStr(s string) string {
	ns := make([]rune, len(s))
	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		ns = append(ns, r)
		s = s[:len(s)-size]
	}
	return string(ns)
}
```

程序输出如下

```shell
reverse before of string: 中国人民解放军8341
reverse after of string: 1438军放解民人国中
```

### 通过rune切片反转字符串

```go
package main

import "fmt"

func main() {
	s := []rune("中国人民解放军8341")
	fmt.Printf("reverse before of string: %s\n", string(s))
	fmt.Printf("reverse after of string: %s\n", ReverseStr(s))
}

// ReverseStr反转字符串
func ReverseStr(p []rune) string {
	ns := make([]rune, len(p))
	for len(p) > 0 {
		var lst rune
		lst, p = p[len(p)-1], p[:len(p)-1]
		ns = append(ns, lst)
	}
	return string(ns)
}
```

程序输出如下

```shell
reverse before of string: 中国人民解放军8341
reverse after of string: 1438军放解民人国中
```

## 分割字符串

### 通过指定索引分割字符串

编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。

```go
package main

import "fmt"

func main() {
	s := "中国人民解放军8341"
	fmt.Printf("split before of string: %s\n", s)

	prefix, suffix := SplitStrByIndex(s, 7)
	fmt.Printf("split after of string, prefix: %s, suffix: %s\n", prefix, suffix)
}

// SplitStrByIndex 通过索引分割字符串
func SplitStrByIndex(str string, index int) (prefix, suffix string) {
	s := []rune(str)
	if index >= 0 && index < len(s) {
		prefix = string(s[:index])
		suffix = string(s[index:])
	}
	return
}
```
程序输出如下

```shell
split before of string: 中国人民解放军8341
split after of string, prefix: 中国人民解放军, suffix: 8341
```

## 对半反转字符串

假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？

```go
package main

import "fmt"

func main() {
	s := "中国人民解放军8341"
	fmt.Printf("reverse before of string: %s\n", s)

	s = ReverseStrByMiddlePosition(s)
	fmt.Printf("reverse after of string: %s\n", s)
}

func ReverseStrByMiddlePosition(s string) string {
	str := []rune(s)
	return string(str[len(str)/2:]) + string(str[:len(str)/2])
}
```

程序输出如下

```shell
reverse before of string: 中国人民解放军8341
reverse after of string: 放军8341中国人民解
```

## 目录
[Back](../GolangNotice.md)