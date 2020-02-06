package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	readString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	// 输出用户输入姓名
	fmt.Printf("Your name is %s\n", readString)

	// 将姓名转换为全小写
	readString = strings.ToLower(readString)
	// For Unix: test with delimiter "\n", for windows: test with "\r\n"
	switch readString {
	case "philip\n", "philip\r\n":
		fmt.Printf("Welcome philip!\n")
	case "chris\n", "chris\r\n":
		fmt.Printf("Welcome chris!\n")
	case "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome ivo!\n")
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}

	// version 2:
	switch readString {
	case "philip\n", "philip\r\n":
		fallthrough
	case "chris\n", "chris\r\n":
		fallthrough
	case "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome %s\n", TrimSpace(readString))
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}

	// version 3:
	switch readString {
	case "philip\n", "philip\r\n", "chris\n", "chris\r\n", "ivo\n", "ivo\r\n":
		fmt.Printf("Welcome %s\n", TrimSpace(readString))
	default:
		fmt.Println("You are not welcome here! GoodBye!")
	}
}

// TrimSpace 去除空格
func TrimSpace(readString string) string {
	return strings.TrimSpace(readString)
}
