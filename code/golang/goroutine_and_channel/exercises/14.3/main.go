package main

import (
	"fmt"
	"time"
)

func main() {
	ints := make(chan int, 100)
	go func() {
		time.Sleep(15e9)
		fmt.Println("Recovered:", <-ints)
	}()
	fmt.Println("Sending:", 10)
	ints <- 10
	fmt.Println("sent:", 10)
}
