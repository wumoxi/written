# Golang『读取用户的输入』注意点

## 使用 Scan 和 Sscan 开头的函数读取用户的输入

```go
package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// 输出结果：Hi 武 沫汐
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	// 输出结果：From the string we read: 56.12 5212 Go
	fmt.Println("From the string we read: ", f, i, s)
}
```

也可以使用 `bufio` 包提供的缓冲读取来读取数据

## 使用缓冲读取数据：示例1

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	readString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The read string was: %s\n", readString)
	}
}
```

这个示例将读取标准输入到数据缓冲，读取到行末时会返回`io.EOF`错误，读取时没有遇到错误时会返回`nil`错误，当不存在错误时`readString`将包含读取到的数据！


## 使用缓冲读取数据：示例2

结合switch分支演示缓冲读取数据！

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	readString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	// 输出用户输入姓名
	fmt.Printf("Your name is %s\n", readString)

	// 将姓名转换为全小写
	readString = strings.ToLower(readString)
	// For Unix: test with delimiter "\n", for windows: test with "\r\n"
	switch readString {
	case "philip\n", "philip\r\n":
		fmt.Printf("Welcome philip!\n")
	case "chris\n", "chris\r\n":
		fmt.Printf("Welcome chris!\n")
	case "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome ivo!\n")
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}

	// version 2:
	switch readString {
	case "philip\n", "philip\r\n":
		fallthrough
	case "chris\n", "chris\r\n":
		fallthrough
	case "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome %s\n", TrimSpace(readString))
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}

	// version 3:
	switch readString {
	case "philip\n", "philip\r\n", "chris\n", "chris\r\n", "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome %s\n", TrimSpace(readString))
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}
}

// TrimSpace 去除空格
func TrimSpace(readString string) string {
	return strings.TrimSpace(readString)
}
```

运行程序输入用户名称为 `ivo`, 程序输出如下：

```shell
Please enter your name:
ivo
Your name is ivo

Welcome ivo!
Welcome ivo
Welcome ivo
```

_**注意**_：Unix和Windows的行结束符是不同的！

## 练习

### 1. 统计字符，单词个数，行数

编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：

1. 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
2. 输入的单词的个数
3. 输入的行数

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop:")

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		if input == "S\r\n" || input == "S\n" {
			fmt.Println("Here are the counts: ")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) - 2
	nrwords += len(strings.Fields(input))
	nrlines++
}
```

运行程序输入相应的内容输出如下：

```shell
Please enter some input, type S to stop:
hello world
my email address is wu.shaohua@foxmail.com
my phone number is 13752223232
my id number is 372328199602238252
my address is 北京昌平区回龙观镇龙乡小区
S
Here are the counts: 
Number of characters: 164
Number of words: 21
Number of lines: 6
```

### 2. 简单逆波兰式计算器

编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束。请使用您在练习11.13中开发的 `stack` 包。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"written/code/golang/read_writer_data/read_user_input/exercises/stack"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	calc := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error!")
			os.Exit(1)
		}

		// 移除"\r\n"
		input = strings.TrimSpace(input)
		switch {
		case input == "q":
			fmt.Println("Calculator stopped!")
			return
		case input >= "0" && input <= "999999":
			i, _ := strconv.Atoi(input)
			calc.Push(i)
		case input == "+":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p+q)
		case input == "-":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p-q)
		case input == "*":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p*q)
		case input == "/":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p/q)
		default:
			fmt.Println("No valid input!")
		}
	}
}
```

运行程序并输入相应的值结果如下：

```shell
Give a number, an operator (+, -, *, /), or q to stop:
3
4
+
The result of 3 + 4 = 7
3
4
-
The result of 3 - 4 = -1
3
4
*
The result of 3 * 4 = 12
3
4
/
The result of 3 / 4 = 0
q
Calculator stopped!
```
