package gofibonacci

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacciNumbers(20, ch)
	start := time.Now()
	index := 0
	for {
		if value, ok := <-ch; ok {
			fmt.Printf("fibonacci(%d) is: %d\n", index, value)
			index++
		} else {
			end := time.Now()
			diff := end.Sub(start)
			fmt.Printf("longCalculation took this amount of time: %s\n", diff)
			os.Exit(0)
		}
	}
}

func fibonacciNumbers(limit int, ch chan int) {
	for i := 0; i <= limit; i++ {
		ch <- fibonacci(i)
	}
	close(ch)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
