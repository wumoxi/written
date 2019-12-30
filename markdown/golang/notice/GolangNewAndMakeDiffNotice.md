# Golang『new和make』区别注意点

## new

new(T) 为每个新的类型T分配一片内存，初始化为T的类型零值并且返回类型为*T的内存地址：这种方法 _**返回一个指向类型为 T，值类型零值的地址的指针**_，它适用于值类型如数组和结构体；它相当于 &T{}。

## make

make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。


## 目录
[Back](../GolangNotice.md)