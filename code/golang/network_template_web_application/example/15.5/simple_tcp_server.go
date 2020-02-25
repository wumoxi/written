package main

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}

	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	addr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address: port failed:'"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", addr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
	}
}

func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())
	}
}
