package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("code/golang/read_writer_data/file_read_writer/input.dat")
	if err != nil {
		fmt.Printf("An error occurred on opening the input file\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	// Close the open file
	defer file.Close()

	// Gets a new reader
	reader := bufio.NewReader(file)

	// Loop reads
	for {
		readString, err := reader.ReadString('\n')
		fmt.Printf("The input was: %s", readString)
		if err != nil && err == io.EOF {
			return
		}
	}
}
