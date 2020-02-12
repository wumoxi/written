package main

import (
	"fmt"
	"os"
)

func main() {
	proAttr := &os.ProcAttr{
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	process, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, proAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", process)
}
