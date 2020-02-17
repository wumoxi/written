package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		fmt.Println("Received:", <-ch)
	}()

	fmt.Println("Sending")
	ch <- 9
	fmt.Println("Sent")
}
