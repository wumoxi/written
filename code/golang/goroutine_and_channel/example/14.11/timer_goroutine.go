package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e9)
	after := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-after:
			fmt.Println("after.")
			return
		default:
			fmt.Println(".")
			time.Sleep(1e9 / 2)
		}
	}
}
