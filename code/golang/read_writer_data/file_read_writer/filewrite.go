package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdout.WriteString("hello world\n")
	file, err := os.OpenFile("code/golang/read_writer_data/file_read_writer/testwrite.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Open the error: %s\n", err)
		return
	}
	defer file.Close()
	file.WriteString("hello, world in a file\n")
}
