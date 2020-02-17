package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, done chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go tel(ch, done)

	for {
		select {
		case v := <-ch:
			fmt.Printf("received: %d\n", v)
		case v := <-done:
			if v {
				fmt.Printf("received done!")
				os.Exit(0)
			}
		}
	}
}
