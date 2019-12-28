# Golang迭代结构for-range注意点


这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。它可以迭代任何一个集合（包括数组和 map以及channel）。语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：`for ix, val := range coll { }`。

要注意的是，val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值（_**译者注：如果 `val` 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值**_）。一个字符串是 Unicode 编码的字符（或称之为 `rune` 它是 int32的别名）集合，因此您也可以用它迭代字符串：

```go
for pos, char := range str {
...
}
```

每个 rune 字符和索引在 for-range 循环中是一一对应的。它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。请看下面的案例

## 使用for-range结构迭代一个字符串

```go
package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is %c\n", pos, char)
	}

	fmt.Println()
	str2 := "Chinese: 汉语😂"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("charcater %c starts at byte position %d\n", char, pos)
	}

	fmt.Println()
	fmt.Println("index\t int(rune)\t rune\t\t char\t bytes")
	for index, char := range str2 {
		fmt.Printf("%-5d\t %-5d\t\t %-7U\t '%c'\t %X\n", index, char, char, char, []byte(string(char)))
	}
}
```

程序具体的输出结果如下所示

```shell
The length of str is: 27
Character on position 0 is G
Character on position 1 is o
Character on position 2 is  
Character on position 3 is i
Character on position 4 is s
Character on position 5 is  
Character on position 6 is a
Character on position 7 is  
Character on position 8 is b
Character on position 9 is e
Character on position 10 is a
Character on position 11 is u
Character on position 12 is t
Character on position 13 is i
Character on position 14 is f
Character on position 15 is u
Character on position 16 is l
Character on position 17 is  
Character on position 18 is l
Character on position 19 is a
Character on position 20 is n
Character on position 21 is g
Character on position 22 is u
Character on position 23 is a
Character on position 24 is g
Character on position 25 is e
Character on position 26 is !

The length of str2 is: 19
charcater C starts at byte position 0
charcater h starts at byte position 1
charcater i starts at byte position 2
charcater n starts at byte position 3
charcater e starts at byte position 4
charcater s starts at byte position 5
charcater e starts at byte position 6
charcater : starts at byte position 7
charcater   starts at byte position 8
charcater 汉 starts at byte position 9
charcater 语 starts at byte position 12
charcater 😂 starts at byte position 15

index	 int(rune)	 rune		 char	 bytes
0    	 67   		 U+0043 	 'C'	 43
1    	 104  		 U+0068 	 'h'	 68
2    	 105  		 U+0069 	 'i'	 69
3    	 110  		 U+006E 	 'n'	 6E
4    	 101  		 U+0065 	 'e'	 65
5    	 115  		 U+0073 	 's'	 73
6    	 101  		 U+0065 	 'e'	 65
7    	 58   		 U+003A 	 ':'	 3A
8    	 32   		 U+0020 	 ' '	 20
9    	 27721		 U+6C49 	 '汉'	 E6B189
12   	 35821		 U+8BED 	 '语'	 E8AFAD
15   	 128514		 U+1F602	 '😂'	 F09F9882
```

我们可以看到，常用英文字符使用 1 个字节表示，而汉字和表情这样的属于多字节字符使用 3 个字节表示。

_**敲黑板请注意了：表情存储到 MYSQL，如果使用`utf8`字符编码进行存储，存储的数据不完整，需要使用`utf8mb4`字符编码才可以完整存储！**_

## 目录
[Back](../../README.md)