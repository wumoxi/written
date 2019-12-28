# Golang『指针』注意点

**_一个指针变量可以指向任何一个值的内存地址_** 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。

当一个指针被定义后没有分配到任何变量时，它的值为 `nil`。

一个指针变量通常缩写为 `ptr`。

## 使用指针减少内存占用和提高效率

指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针传递是很廉价的(_**指针拷贝**_)，只占用 4 个或 8 个字节。当程序在工作中需要占用大量的内存，或很多变量，或者两者都有，使用指针会减少内存占用和提高效率。被指向的变量也保存在内存中，直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。

## 关于 Go 语言中 * 号的所有用法

- `var p *int`: 表示一个指针变量。
- `*p`: 表示获取指针变量`p`指向地址上所存储的值。

## 目录
[Back](../GolangNotice.md)