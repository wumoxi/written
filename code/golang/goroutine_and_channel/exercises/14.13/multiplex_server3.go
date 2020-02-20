package main

import (
	"fmt"
)

type Request struct {
	a     int
	b     int
	reply chan int
}

func (r *Request) String() string {
	return fmt.Sprintf("%d + %d = %d", r.a, r.b, <-r.reply)
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

	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}

	adder <- req1
	adder <- req2
	fmt.Println(req1, req2)

	// notice to the channel of the 'done'
	quit <- true

	fmt.Println("All request is process done!")
}
