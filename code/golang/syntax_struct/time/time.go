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
