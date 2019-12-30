# Golang『切片』注意点

## 切片的复制与追加

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。下面的代码描述了从拷贝切片的 copy 函数和向切片追加新元素的 append 函数。

```go
package main
import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
```

`func append(s[]T, x ...T) []T` 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素同类型。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。

如果你想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可：`x = append(x, y...)`。

**注意**： append 在大多数情况下很好用，但是如果你想完全掌控整个追加过程，你可以实现一个这样的 AppendByte 方法：

```go
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
```

`func copy(dst, src []T) int` copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数。源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。如果 src 是字符串那么元素类型就是 byte。如果你还想继续使用 src，在拷贝结束后执行 `src = dst`。

## 扩展切片的长度

给定一个slice`s []int` 和一个 int 类型的因子factor，扩展 s 使其长度为 `len(s) * factor`。

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("value: %+v, length:%d, capacity: %d\n", s, len(s), cap(s))
	s = extendedSliceLength(s, 2)
	fmt.Printf("value: %+v, length:%d, capacity: %d\n", s, len(s), cap(s))
	s = extendedSliceLength(s, 4)
	fmt.Printf("value: %+v, length:%d, capacity: %d\n", s, len(s), cap(s))
}

func extendedSliceLength(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	s = ns
	return s
}
```

程序输出如下

```shell
value: [1 2 3], length:3, capacity: 3
value: [1 2 3 0 0 0], length:6, capacity: 6
value: [1 2 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], length:24, capacity: 24
```


## 过滤切片元素

用回调函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 `fn func(int) bool`，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。

```shell
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Before processing: %+v\n", s)
	s = Filter(s, func(number int) bool {
		return number%2 == 0
	})
	fmt.Printf("After processing: %+v\n", s)
}

func Filter(s []int, fn func(int) bool) (result []int) {
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			result = append(result, s[i])
		}
	}
	return
}
```

程序输出如下

```shell
Before processing: [1 2 3 4 5 6 7 8 9 10]
After processing: [2 4 6 8 10]
```

## 插入切片到另一个切片的指定索引位置

写一个函数 InsertStringSlice 将切片插入到另一个切片的指定索引位置。

```go
package main

import "fmt"

func main() {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	d := []string{"8859", "8868"}
	s = InsertStringSlice(s, d, len(s)-1)
	fmt.Printf("Insert after of slice: %v, length: %d\n", s, len(s))
	s = InsertStringSlice(s, d, 0)
	fmt.Printf("Insert after of slice: %v, length: %d\n", s, len(s))
	s = InsertStringSlice(s, d, -1)
	fmt.Printf("Insert after of slice: %v, length: %d\n", s, len(s))
}

// InsertStringSlice插入切片到另一个切片指定索引位置
func InsertStringSlice(dst []string, data []string, index int) []string {
	if index >= 0 && index < len(dst) {
		dl := len(data)
		nl := len(dst) + dl
		n := make([]string, nl)

		prefix := dst[:index]
		pl := len(prefix)

		suffix := dst[index:]

		// 生成切片元素
		copy(n[0:pl], prefix)
		copy(n[pl:pl+dl], data)
		copy(n[pl+dl:], suffix)

		return n
	}
	return dst
}
```

程序输出结果如下

```shell
Insert after of slice: [1 2 3 4 5 6 7 8 9 8859 8868 10], length: 12
Insert after of slice: [8859 8868 1 2 3 4 5 6 7 8 9 8859 8868 10], length: 14
Insert after of slice: [8859 8868 1 2 3 4 5 6 7 8 9 8859 8868 10], length: 14
```

## 删除切片元素

写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。

```go
package main

import "fmt"

func main() {
	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	fmt.Printf("Remove before of slice: %v\n", s)
	fmt.Printf("Remove after of slice: %v\n", RemoveStringSlice(s, 1, 2))
}

// RemoveStringSlice删除切片元素
func RemoveStringSlice(src []string, start, end int) []string {
	if (start < end) && (start >= 0 && start < len(src) && end < len(src)) {
		src = append(src[:start], src[end:]...)
	}
	return src
}
```


程序输出结果如下

```shell
Remove before of slice: [1 2 3 4 5 6 7 8 9 10]
Remove after of slice: [1 3 4 5 6 7 8 9 10]
```

## 目录
[Back](../GolangNotice.md)