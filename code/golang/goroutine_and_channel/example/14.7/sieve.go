package main

import "fmt"

func generate(ch chan int, limit int) {
	for i := 2; i < limit; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in chan int, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	send := make(chan int)
	go generate(send, 100)
	for {
		prime := <-send
		fmt.Print(prime, " ")
		receiver := make(chan int)
		go filter(send, receiver, prime)
		send = receiver
	}
}
