# Golang『日期和时间』注意点

## Golang 特定的时间格式化字符串


- `2006`: 年 
- `01`: 月
- `02`: 日
- `03`: 小时
- `04`: 分钟
- `05`: 秒

### 要格式化当前时间

例如，要格式化当前时间，为中国通用时间格式，应该这么做

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format("2006-01-02 03:04:05")
	fmt.Printf("当前时间: %s\n", now)
}
```

程序输出：

```shell
当前时间: 2019-12-27 10:54:28
```

### 格式化一个特定的时间戳

例如，要格式化一个特定的时间戳，为中国通用时间格式，应该这么做

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 过去
	formerly := time.Now().Add(-(60 * 60 * 24 * 7 * 1e9)).Unix()
	formerlyTime := time.Unix(formerly, 0).Format("2006-01-02 03:04:05")
	fmt.Printf("过去7天的时间: %s\n", formerlyTime)

	// 未来
	future := time.Now().Add(60 * 60 * 24 * 7 * 1e9).Unix()
	futureTime := time.Unix(future, 0).Format("2006-01-02 03:04:05")
	fmt.Printf("未来7天的时间: %s\n", futureTime)
}
```

程序输出：

```shell
过去7天的时间: 2019-12-20 10:54:28
未来7天的时间: 2020-01-03 10:54:28
```

## 程序执行一次或多次

如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 `time.After` 或者 `time.Ticker` 进行处理，它们都会返回一个时间管道值。
 
## 对某个进程或者说是Goroutine暂停一个时间段 

`time.Sleep(d Duration)` 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停。

## 目录
[Back](../GolangNotice.md)