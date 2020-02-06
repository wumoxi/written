package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"written/code/golang/read_writer_data/read_user_input/exercises/stack"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	calc := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error!")
			os.Exit(1)
		}

		// 移除"\r\n"
		input = strings.TrimSpace(input)
		switch {
		case input == "q":
			fmt.Println("Calculator stopped!")
			return
		case input >= "0" && input <= "999999":
			i, _ := strconv.Atoi(input)
			calc.Push(i)
		case input == "+":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p+q)
		case input == "-":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p-q)
		case input == "*":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p*q)
		case input == "/":
			q := calc.Pop()
			p := calc.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, input, q, p/q)
		default:
			fmt.Println("No valid input!")
		}
	}
}
