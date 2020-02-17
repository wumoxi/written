package main

import "fmt"

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go tel(ch)
	for v := range ch {
		fmt.Printf("received: %d\n", v)
	}
}
