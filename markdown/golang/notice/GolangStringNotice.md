# Golang『字符串』注意点

## 反转字符串

```go
// ReverseStr反转字符串
func ReverseStr(p []byte) string {
	s := make([]rune, len(p))
	for len(p) > 0 {
		r, size := utf8.DecodeLastRune(p)
		s = append(s, r)
		p = p[:len(p)-size]
	}
	return string(s)
}
```