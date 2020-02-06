package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	readString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The read string was: %s\n", readString)
	}
}
