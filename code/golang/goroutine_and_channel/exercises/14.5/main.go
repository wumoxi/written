package main

import "fmt"

func main() {
	ints := make(chan int)
	go sum(10, 8, ints)
	fmt.Println(<-ints)
}

func sum(a int, b int, ints chan int) {
	ints <- a + b
}
