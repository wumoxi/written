# Golang『闭包』注意点

## 斐波那契数

### 使用闭包获取指定数字的斐波那契数

```go
package main

import "fmt"

func main() {
	fb := fib()
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%02d) is: %d\n", i, fb())
	}
}

// fib获取斐波那契数
func fib() func() int {
	a, b := 1, 1
	return func() int {
		c := a
		a, b = b, b+c
		return c
	}
}
```
程序输出0到10之间的斐波那契数如下

```shell
fib(00) is: 1
fib(01) is: 1
fib(02) is: 2
fib(03) is: 3
fib(04) is: 5
fib(05) is: 8
fib(06) is: 13
fib(07) is: 21
fib(08) is: 34
fib(09) is: 55
fib(10) is: 89
```


## 目录
[Back](../GolangNotice.md)