# Golang迭代结构for注意点

## 以下程序的输出结果是什么？

```go
for i := 0; i < 5; i++ {
	var v int
	fmt.Printf("%d ", v)
	v = 5
}
```

它会输出 `0 0 0 0 0`，而不是 `5 5 5 5 5`，具体原因也很简单，变量 `v` 在输出前只是声明而没有初始化，那么对于整型`int`来讲，只声明而没有初始化，它的值就是默认零值，所以就是整数`0`。

## 请描述以下 for 循环的输出结果？

```go
for i := 0; ; i++ {
	fmt.Println("Value of i is now:", i)
}
```

它是一个无限循环，会输出类似的 `Value of i is now：0`、`Value of i is now：1`...`Value of i is now：n`，只要你不手动终止循环，它就不会终止！

## 请描述以下 for 循环的输出结果？

```go
for i := 0; i < 3; {
	fmt.Println("Value of i:", i)
}
```

它是一个无限循环，会输出`Value of i: 0`，无限次，因为变量`i`的值永远是`0`，所以条件判断会永远成立，只要你不手动终止循环，它就不会终止！

## 请描述以下 for 循环的输出结果？

```go
s := ""
for ; s != "aaaaa"; {
	fmt.Println("Value of s:", s)
	s = s + "a"
}
```

它会输出如下结果

```shell
Value of s: 
Value of s: a
Value of s: aa
Value of s: aaa
Value of s: aaaa
```

一共输出5次，当将要输出第6次的时候，变量 `s` 的值为 `aaaaa`，此时条件判断结果为`false`，这时循环终止！

## 请描述以下 for 循环的输出结果？

```go
for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
	fmt.Println("Value of i, j, s:", i, j, s)
}
```

它会输出如下结果

```shell
Value of i, j, s: 0, 5, "a"
Value of i, j, s: 1, 6, "aa"
Value of i, j, s: 2, 7, "aaa"
```

好了，这是一个很经典的多条件判断for循环，这就是最最基础的语法考察与应用，判断条件中有三个表达式分别是 `i < 3`、`j < 100`、`s != "aaaaa"`，其中决定循环次数的条件表达式就是 `i < 3`，这个for循环一共会循环3次，原因也很简单嘛，变量`i`的初始化值为`0`，当执行到第四次的时候，它的值变为3，所以判断条件不成立，循环也就结束了！

_**注意了：如果在真正的项目中这么写是没有意思的，这只在面试或学习中才会这么写，知识点考察嘛！**_


## 目录
[Back](../../README.md)
