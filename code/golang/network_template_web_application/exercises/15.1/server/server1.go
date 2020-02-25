package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	mapUsers = make(map[string]int)
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
		_, err := conn.Read(buff)
		checkError(err)
		input := strings.TrimSpace(string(buff))
		if strings.Contains(input, "@SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}

		if strings.Contains(input, "@WHO") {
			DisplayList()
		}

		clName := input[0 : strings.Index(input, "says")-1]
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--\n", string(buff))
	}
}

func DisplayList() {
	fmt.Println("---------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("---------------------------------------------")
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}
