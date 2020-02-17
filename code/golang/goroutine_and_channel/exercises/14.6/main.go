package main

import "fmt"

func main() {
	ints := make(chan int)
	bools := make(chan bool)
	go generate(0, 10, ints)
	go read(ints, bools)
	<-bools
}

func read(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func generate(start int, count int, out chan<- int) {
	for i := start; i < count; i++ {
		out <- start
		start += count
	}
	close(out)
}
