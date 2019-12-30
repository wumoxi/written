# Golang『new和make』区别注意点

## new

new(T) 为每个新的类型T分配一片内存，初始化为T的类型零值并且返回类型为*T的内存地址：这种方法 _**返回一个指向类型为 T，值为类型零值的地址的指针**_，它适用于值类型如数组和结构体；它相当于 &T{}。

## make

make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。

**换言之，new 函数分配内存，make 函数初始化；**

## 如何理解new、make、slice、map、channel的关系

1.slice、map以及channel都是golang内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而make就是对三者进行初始化的一种操作方式

2.new 为指定数值类型分配内存，并返回指向类型地址的指针，只分配内存不初始化， 所以slice、map、channel需要make进行初始化并获取对应的内存地址，而非new简单的内存分配获取内存地址

## 目录
[Back](../GolangNotice.md)