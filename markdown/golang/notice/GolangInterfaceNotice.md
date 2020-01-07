# Golang『接口』注意点

## Go的接口相对其它解析性语言有什么特点？

Go 的接口提高了代码的分离度，改善了代码的复用性，使得代码开发过程中的设计模式更容易实现。用 Go 接口还能实现 `依赖注入模式`。

## 泛型

通过空接口和类型断言，现在我们可以写一个可以应用于许多类型的 泛型 的 map 函数，为 int 和 string 构建一个把 int 值加倍和将字符串值与其自身连接（译者注：即"abc"变成"abcabc"）的 map 函数 mapFunc。为了可读性可以定义一个 interface{} 的别名，比如：type obj interface{}。

```go
package main

import (
	"fmt"
	"strings"
)

type o interface {}

func main() {
	// 定义通用lambda(匿名)函数mf
	mf := func(i o) o {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return strings.Repeat(i.(string), 2)
		case bool:
			if i.(bool) {
				return "真"
			}
			return "假"
		}
		return i
	}

	// 整型切片
	iList := []o{1, 2, 3, 4, 5, 6}
	irList := mapFunc(mf, iList)
	for _, v := range irList {
		fmt.Println(v)
	}

	// 输出空白行
	fmt.Println()

	// 字符串切片
	sList := []o{"1", "2", "3", "4", "5", "6"}
	srList := mapFunc(mf, sList)
	for _, v := range srList {
		fmt.Println(v)
	}

	// 输出空白行
	fmt.Println()

	// 多类型切片
	mList := []o{"hello", 33.5, true, 22, 'B'}
	mrList := mapFunc(mf, mList)
	for _, v := range mrList {
		fmt.Println(v)
	}
}

func mapFunc(mf func(o) o, list []o) []o {
	r := make([]o, len(list))
	for ix, v := range list {
		r[ix] = mf(v)
	}
	return r
}
```

程序输出如下

```shell
2
4
6
8
10
12

11
22
33
44
55
66

hellohello
33.5
真
44
66
```

## 可变参

稍微改变上面的泛型示例，允许 mapFunc 接收不定数量的 items。

```go

package main

import (
	"fmt"
	"strings"
)

type o interface {}

func main() {
	// 定义通用lambda(匿名)函数mf
	mf := func(i o) o {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return strings.Repeat(i.(string), 2)
		case bool:
			if i.(bool) {
				return "真"
			}
			return "假"
		}
		return i
	}

	// 整型切片
	irList := mapFunc(mf, 1, 2, 3, 4, 5, 6)
	for _, v := range irList {
		fmt.Println(v)
	}

	// 输出空白行
	fmt.Println()

	// 字符串切片
	srList := mapFunc(mf, "1", "2", "3", "4", "5", "6")
	for _, v := range srList {
		fmt.Println(v)
	}

	// 输出空白行
	fmt.Println()

	// 多类型切片
	mrList := mapFunc(mf, "hello", 33.5, true, 22, 'B')
	for _, v := range mrList {
		fmt.Println(v)
	}
}

func mapFunc(mf func(o) o, list ...o) []o {
	r := make([]o, len(list))
	for ix, v := range list {
		r[ix] = mf(v)
	}
	return r
}
```

程序输出如下

```shell
2
4
6
8
10
12

11
22
33
44
55
66

hellohello
33.5
真
44
66
```

### 接口实现栈

```go
package main

import (
	"errors"
	"fmt"
)

type Stacker interface {
	Len() int
	IsEmpty() bool
	Push(interface{})
	Pop() (interface{}, error)
}

type Stack struct {
	Id     int           // 索引表示第一个空闲的位置
	Bucket []interface{} // 栈数据桶
}

func (s *Stack) Len() int {
	return len(s.Bucket)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(item interface{}) {
	s.Bucket = append(s.Bucket, item)
}

func (s *Stack) Pop() (item interface{}, err error) {
	if !s.IsEmpty() {
		item = s.Bucket[len(s.Bucket)-1]
		s.Bucket = s.Bucket[:len(s.Bucket)-1]

		return item, nil
	}
	return item, errors.New("stack is empty")
}

func (s *Stack) String() (str string) {
	if !s.IsEmpty() {
		for i, v := range s.Bucket {
			if len(str) != 0 {
				str += " "
			}
			str += fmt.Sprintf("[%v:%v]", i, v)
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
	stack.Push(true)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(3)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(4)
	fmt.Printf("入栈后栈数据：%s\n", stack)
	stack.Push(5)
	fmt.Printf("入栈后栈数据：%s\n\n", stack)

	// 出栈
	item, err := stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
	item, err = stack.Pop()
	fmt.Printf("弹出栈：%v\t, 错误: %v, 出栈后栈数据：%s\n", item, err, stack)
}
```

程序输出如下

```shell
入栈后栈数据：[0:1]
入栈后栈数据：[0:1] [1:true]
入栈后栈数据：[0:1] [1:true] [2:3]
入栈后栈数据：[0:1] [1:true] [2:3] [3:4]
入栈后栈数据：[0:1] [1:true] [2:3] [3:4] [4:5]

弹出栈：5	, 错误: <nil>, 出栈后栈数据：[0:1] [1:true] [2:3] [3:4]
弹出栈：4	, 错误: <nil>, 出栈后栈数据：[0:1] [1:true] [2:3]
弹出栈：3	, 错误: <nil>, 出栈后栈数据：[0:1] [1:true]
弹出栈：true	, 错误: <nil>, 出栈后栈数据：[0:1]
弹出栈：1	, 错误: <nil>, 出栈后栈数据：[]
弹出栈：<nil>	, 错误: stack is empty, 出栈后栈数据：[]
弹出栈：<nil>	, 错误: stack is empty, 出栈后栈数据：[]
```