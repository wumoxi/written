# Golang『自定义包中的错误处理和panicking』注意点

这是所有自定义包实现者应该遵守的最佳实践：

1）_**在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()**_

2）_**向包的调用者返回错误值（而不是 panic）。**_

在包内部，特别是在非导出函数中有很深层次的嵌套调用时，将 panic 转换成 error 来告诉调用方为何出错，是很实用的（且提高了代码可读性）。

下面的示例程序则很好地阐述了这一点。我们有一个简单的 parse 包（[示例 13.4](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_13/parse/parse.go)）用来把输入的字符串解析为整数切片；这个包有自己特殊的 `Error`。

示例：parse.go

```go
package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// Error indicates an error in converting a word into an integer. (指示将单词转换为整数时出错)
type Error struct {
	Index int    // The index into the space-separated list of words. (空格分隔的单词列表的索引)
	Word  string // The word that generated the parse error. (生成解析错误的单词)
	Err   error  // The raw error that precipitated this error, if any. (引发此错误的原始错误(如果有的话))
}

// String returns a human-readable error message. (返回人类可读的错误消息)
func (e *Error) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in input as integers.(将输入中的空格分隔的单词解析为整数)
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2number(fields)
	return
}

func fields2number(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&Error{Index: idx, Word: field, Err: err})
		}
		numbers = append(numbers, num)
	}
	return
}
```

当没有东西需要转换或者转换成整数失败时，这个包会 panic（在函数 fields2numbers 中）。但是可导出的 Parse 函数会从 panic 中 recover 并用所有这些信息返回一个错误给调用者。为了演示这个过程，请看下面的示例程序；不可解析的字符串会导致错误并被打印出来。

```go
package main

import (
	"fmt"
	"written/code/golang/error_process_and_test/custom_package_error_process/parse"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"lst class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n ", ex)
		numbers, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(numbers)
	}
}
```

运行程序将输出如下所示信息：

```shell
Parsing "1 2 3 4 5":
 [1 2 3 4 5]
Parsing "100 50 25 12.5 6.25":
 pkg: pkg parse: error parsing "12.5" as int
Parsing "2 + 2 = 4":
 pkg: pkg parse: error parsing "+" as int
Parsing "lst class":
 pkg: pkg parse: error parsing "lst" as int
Parsing "":
 pkg: no words to parse
```

## 目录
[Back](../GolangNotice.md)