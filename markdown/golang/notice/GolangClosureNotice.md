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

## 闭包工厂函数

一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。下面的函数演示了如何动态返回追加后缀的函数：

```go
// MakeAddSuffix闭包工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}
```

现在，我们可以生成如下函数：

```go
addGif := MakeAddSuffix(".gif")
addPng := MakeAddSuffix(".png")
```

然后调用它们：

```go
gif := addGif("avatar") // avatar.gif
png := addPng("avatar") // avatar.png
```

可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。很显然 Go 语言具有一些函数式语言的特性。闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作。

## 目录
[Back](../GolangNotice.md)