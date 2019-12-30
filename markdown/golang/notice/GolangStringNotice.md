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


## 目录
[Back](../GolangNotice.md)