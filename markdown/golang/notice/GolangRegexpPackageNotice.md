# Golang『regexp package』注意点

## 简单模式

如果是简单模式，使用 `Match` 方法便可：

```go
ok, _ := regexp.Match(pat, []byte(searchIn))
```

变量 ok 将返回 true 或者 false,我们也可以使用 `MatchString`：

```go
ok, _ := regexp.MatchString(pat, searchIn)
```

_**以下示例为简单模式使用案例**_

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Printf("is matched: %t\n", IsMobile("13817995589"))
	fmt.Printf("is matched: %t\n", IsMobile("+86-13817995589"))
	fmt.Printf("is matched: %t\n", IsMobile("-13817995589"))
}

// IsMobile检测是否是一个有效的手机号
func IsMobile(mobile string) bool {
	const patten = `^[\+|1](\d{2}-1)?[3-9]{1}\d{9}`
	// matched, _ := regexp.Match(patten, []byte(mobile))
	matched, _ := regexp.MatchString(patten, mobile)
	return matched
}
```

程序输出如下

```shell
is matched: true
is matched: true
is matched: false
```

## 复杂模式

更多方法中，必须先将正则模式通过 Compile 或 MustCompile 方法返回一个 Regexp 对象。然后我们将掌握一些匹配，查找，替换相关的功能。

_**以下为复杂模式使用案例**_

```go
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`[0-9]+.[0-9]+`)

func main() {
	// 目标字符串
	str := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	fmt.Printf("original string: %s\n", str)

	// ReplaceAllString
	fmt.Printf("replaced string: %s\n", re.ReplaceAllString(str, "##.#"))

	// ReplaceAllStringFunc
	fmt.Printf("replaced string: %s\n", re.ReplaceAllStringFunc(str, func(s string) string {
		return strconv.FormatFloat(parseFloatByStr(s)*2, 'f', 2, 64)
	}))

	// FindAllString
	fmt.Printf("finds: %+v\n", SalariesByStr(re, str))
}

// parseFloatByStr将字符串解析为float64
func parseFloatByStr(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// SalariesByStr 获取薪水数据
func SalariesByStr(re *regexp.Regexp, str string) (salaries []float64) {
	all := re.FindAllString(str, -1)
	for _, item := range all {
		salaries = append(salaries, parseFloatByStr(item))
	}
	return
}
```

程序输出如下

```shell
original string: John: 2578.34 William: 4567.23 Steve: 5632.18
replaced string: John: ##.# William: ##.# Steve: ##.#
replaced string: John: 5156.68 William: 9134.46 Steve: 11264.36
finds: [2578.34 4567.23 5632.18]
```

`Compile` 函数也可能返回一个错误，我们在使用时忽略对错误的判断是因为我们确信自己正则表达式是有效的。当用户输入或从数据中获取正则表达式的时候，我们有必要去检验它的正确性。另外我们也可以使用 `MustCompile` 方法，它可以像 `Compile` 方法一样检验正则的有效性，但是当正则不合法时程序将 `panic`。

## 目录
[Back](../GolangNotice.md)