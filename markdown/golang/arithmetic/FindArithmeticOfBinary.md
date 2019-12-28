# 二分查找

```go
// 二分查找递归, 时间复杂度为 O(log2n), 空间复杂度：O(log2n)
package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	l := 0
	h := len(s) - 1
	f := 8
	fmt.Printf("find %d, get position of index: %d\n", f, BinaryFind(s, f, l, h))
}

// BinaryFind 二分查找，查找到元素的索引位置，存在则返回具体索引，否则返回-1
func BinaryFind(arr []int, search, low, height int) int {
	// 查找
	if low <= height {
		// 获取中间数索引值
		m := int(math.Floor(float64((low + height) / 2)))
		// 查找元素
		switch {
		case arr[m] == search:
			return m
		case arr[m] > search:
			return BinaryFind(arr, search, low, m+1)
		case arr[m] < search:
			return BinaryFind(arr, search, m+1, height)
		}
	}
	return -1
}
```

## 目录
[Back](../GolangArithmetic.md)