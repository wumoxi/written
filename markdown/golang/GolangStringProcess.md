# Golang字符串处理

## 字符串翻转

这个处理方式处理很简单也很暴力！

```go
// StrReverse翻转字符串
func StrReverse(str string) string {
	letters := strings.Split(str, "")
	res := make([]string, len(letters))
	for i := len(letters) - 1; i >= 0; i-- {
		res = append(res, letters[i])
	}
	return strings.Join(res, "")
}
```