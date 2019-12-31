# Golang『unsafe package』注意点

## 获取变量占用字节数

```go
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var i int = 110
	fmt.Printf("size of int: %d\n", unsafe.Sizeof(i))

	var i32 int32 = math.MaxInt32
	fmt.Printf("size of int32: %d\n", unsafe.Sizeof(i32))

	var i64 int64 = math.MaxInt64
	fmt.Printf("size of int64: %d\n", unsafe.Sizeof(i64))
}
```

程序输出如下

```shell
size of int: 8
size of int32: 4
size of int64: 8
```

## 目录
[Back](../GolangNotice.md)