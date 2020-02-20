package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Starting the server...")

	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:8859")
	if err != nil {
		fmt.Println("Error listening", err)
		// 终止程序
		return
	}

	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err)
			continue
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buff := make([]byte, 512)
		len, err := conn.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading", err)
			}
			return
		}
		fmt.Printf("Received data: %v\n", string(buff[:len]))
	}
}
