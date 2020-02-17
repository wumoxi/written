package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		// Loop read data for the channel
		for i := 0; i < 30; i++ {
			fmt.Printf("fibonacci(%d) is: %d\n", i, <-ch)
		}

		// Write data to the channel
		done <- true
	}()
	fibonacci(ch, done)
	end := time.Now()
	fmt.Printf("calc time: %v", end.Sub(start))
}

func fibonacci(c chan int, done chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // send data to the channel
			x, y = y, x+y
		case <-done: // read data for the channel
			fmt.Println("Send data to the channel is done!")
			return
		}
	}
}
