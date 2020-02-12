package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("ls", "-l")
	err := command.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", command)
}
