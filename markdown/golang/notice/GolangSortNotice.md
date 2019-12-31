# Golang『排序』注意点

## 冒泡排序

```go
package main

import (
	"fmt"
)

func main() {
	s := []int{8, 2, 3, 5, 6, 1, 9}
	fmt.Printf("sorted before of string: %v\n", s)

	fmt.Printf("sorted after of string: %v\n", BubbleSort(s))
}

// BubbleSort冒泡排序
func BubbleSort(s []int) []int {
	l := len(s) - 1
	for i := 0; i < l; i++ {
		for j := 0; j < l-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
```

程序输出如下

```shell
reverse before of string: [8 2 3 5 6 1 9]
reverse after of string: [1 2 3 5 6 8 9]
```