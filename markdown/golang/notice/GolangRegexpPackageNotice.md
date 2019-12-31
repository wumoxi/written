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

以下示例为简单模式使用案例

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

## 目录
[Back](../GolangNotice.md)