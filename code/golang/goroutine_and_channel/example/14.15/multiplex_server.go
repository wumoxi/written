package main

import (
	"fmt"
)

type Request struct {
	a     int
	b     int
	reply chan int
}

type binOp func(a, b int) int

func run(op binOp, request *Request) {
	request.reply <- op(request.a, request.b)
}

func server(op binOp, services chan *Request, quit chan bool) {
	for {
		select {
		case s := <-services:
			go run(op, s)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (chan *Request, chan bool) {
	service := make(chan *Request)
	quit := make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer(func(a, b int) int {
		return a + b
	})
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.reply = make(chan int)
		adder <- req
	}

	// Checks
	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].reply != 2*i+N {
			fmt.Printf("Fail at: %d\n", i)
		} else {
			fmt.Printf("Request %d is done!\n", i)
		}
	}

	// notice to the channel of the 'done'
	quit <- true

	fmt.Println("All request is process done!")
}
