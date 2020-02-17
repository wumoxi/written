package main

import "fmt"

func main() {
	ints := make(chan int, 100)
	fmt.Println(cap(ints)-len(ints), len(ints))
	ints <- 100
	fmt.Println(cap(ints)-len(ints), len(ints))
	chanints := make(chan int)
	fmt.Println(cap(chanints))
}
