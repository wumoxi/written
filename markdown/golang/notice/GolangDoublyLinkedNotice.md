# Golang『双向链表』注意点

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个新列表，并在其中添加一些数字或字符串。
	lst := list.New()
	lst.PushBack(101)
	lst.PushBack(102)
	lst.PushBack(103)
	lst.PushBack("hello world")

	// 遍历列表并打印其内容(向后遍历)。
	// list.Front() 获取链表头部指针
	for element := lst.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// 遍历列表并打印其内容(向前遍历)。
	// list.Back() 获取链表尾部指针
	for element := lst.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
```

程序输出如下

```shell
101
102
103
hello world
hello world
103
102
101
```

## 目录
[Back](../GolangNotice.md)