package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const delim = '\n'

func main() {
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:8859")
	if err != nil {
		// 由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	name, err := reader.ReadString(delim)
	if err != nil {
		fmt.Println("Error ReadString", err)
		return
	}
	name = strings.TrimSpace(name)

	// 给服务器发送信息直到程序退出
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, err := reader.ReadString(delim)
		if err != nil {
			fmt.Println("Error ReadString", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "Q" {
			return
		}
		_, err = conn.Write([]byte(name + " says: " + input))
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error Writing", err)
			}
			return
		}
	}
}
