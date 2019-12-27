package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(StrReverse("hello, world! 中国"))
}

// StrReverse翻转字符串
func StrReverse(str string) string {
	letters := strings.Split(str, "")
	res := make([]string, len(letters))
	for i := len(letters) - 1; i >= 0; i-- {
		res = append(res, letters[i])
	}
	return strings.Join(res, "")
}
