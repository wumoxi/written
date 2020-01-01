# Golang『unsafe package』注意点

## 获取变量占用字节数

### 标准类型

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

### 结构体

```go
package main

import (
	"fmt"
	"unsafe"
)

type Address struct {
	string
}

type VCard struct {
	Name string
	Address []*Address
	Birth string
	Avatar string
}


func main() {
	vcard := VCard{Name:"张三", Address:[]*Address{&Address{"北京"}, &Address{"上海"}}, Birth:"2019", Avatar:"avatar.png"}

	fmt.Printf("结构体值占用内存: %d\n", unsafe.Sizeof(vcard))

	fmt.Printf("结构体指针占用内存: %d\n", unsafe.Sizeof(&vcard))
}
```

程序输出如下

```shell
结构体值占用内存: 72
结构体指针占用内存: 8
```

## 目录
[Back](../GolangNotice.md)