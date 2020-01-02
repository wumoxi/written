# Golang『栈』简单实现

## 切片实现


```go
package main

import "fmt"

type Stack []int

func (s *Stack)Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (item int, exist bool) {
	if len(*s) > 0 {
		item = (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return item, true
	}
	return item, false
}

func (s *Stack)String() (str string) {
	if len(*s) > 0 {
		for i, v := range *s {
			if len(str) != 0 {
				str += " "
			}
			str += fmt.Sprintf("[%d:%d]", i, v)
		}
	} else {
		str = "[]"
	}
	return
}

func main() {
	// 初始化栈实例
	stack := new(Stack)

	// 入栈
	stack.Push(1)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(2)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(3)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(4)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(5)
	fmt.Printf("入栈后栈数据：%s\n\n", stack)

	// 出栈
	item, exist := stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
}
```

程序输出如下

```shell
入栈后栈数据：[0:1]
入栈后栈数据：[0:1] [1:2]
入栈后栈数据：[0:1] [1:2] [2:3]
入栈后栈数据：[0:1] [1:2] [2:3] [3:4]
入栈后栈数据：[0:1] [1:2] [2:3] [3:4] [4:5]

弹出栈：5, 是否存在: true, 出栈后栈数据：[0:1] [1:2] [2:3] [3:4]
弹出栈：4, 是否存在: true, 出栈后栈数据：[0:1] [1:2] [2:3]
弹出栈：3, 是否存在: true, 出栈后栈数据：[0:1] [1:2]
弹出栈：2, 是否存在: true, 出栈后栈数据：[0:1]
弹出栈：1, 是否存在: true, 出栈后栈数据：[]
弹出栈：0, 是否存在: false, 出栈后栈数据：[]
弹出栈：0, 是否存在: false, 出栈后栈数据：[]
```

## 结构体实现

```go
package main

import "fmt"

type Stack struct {
	Id     int   // 索引表示第一个空闲的位置
	Bucket []int // 栈数据桶
}

func (s *Stack) Push(val int) {
	s.Bucket = append(s.Bucket, val)
}


func (s *Stack) Pop() (item int, exist bool) {
	if len(s.Bucket) > 0 {
		this := &(*s)

		item = this.Bucket[len(this.Bucket)-1]
		this.Bucket = this.Bucket[:len(this.Bucket)-1]

		return item, true
	}
	return item, false
}

func (s *Stack)String() (str string) {
	if len(s.Bucket) > 0 {
		for i, v := range s.Bucket {
			if len(str) != 0 {
				str += " "
			}
			str += fmt.Sprintf("[%d:%d]", i, v)
		}
	} else {
		str = "[]"
	}
	return
}

func main() {
	// 初始化栈实例
	stack := new(Stack)

	// 入栈
	stack.Push(1)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(2)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(3)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(4)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(5)
	fmt.Printf("入栈后栈数据：%s\n\n", stack)

	// 出栈
	item, exist := stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
}
```

程序输出如下

```shell
package main

import "fmt"

type Stack struct {
	Id     int   // 索引表示第一个空闲的位置
	Bucket []int // 栈数据桶
}

func (s *Stack) Push(val int) {
	s.Bucket = append(s.Bucket, val)
}


func (s *Stack) Pop() (item int, exist bool) {
	if len(s.Bucket) > 0 {
		this := &(*s)

		item = this.Bucket[len(this.Bucket)-1]
		this.Bucket = this.Bucket[:len(this.Bucket)-1]

		return item, true
	}
	return item, false
}

func (s *Stack)String() (str string) {
	if len(s.Bucket) > 0 {
		for i, v := range s.Bucket {
			if len(str) != 0 {
				str += " "
			}
			str += fmt.Sprintf("[%d:%d]", i, v)
		}
	} else {
		str = "[]"
	}
	return
}

func main() {
	// 初始化栈实例
	stack := new(Stack)

	// 入栈
	stack.Push(1)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(2)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(3)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(4)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(5)
	fmt.Printf("入栈后栈数据：%s\n\n", stack)

	// 出栈
	item, exist := stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
	item, exist = stack.Pop()
	fmt.Printf("弹出栈：%d, 是否存在: %t, 出栈后栈数据：%s\n", item, exist, stack)
}
```
