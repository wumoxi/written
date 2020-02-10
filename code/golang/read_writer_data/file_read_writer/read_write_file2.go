package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// Data source and destination
	inputFile := "code/golang/read_writer_data/file_read_writer/description.txt"
	outputFile := "code/golang/read_writer_data/file_read_writer/description_copy.txt"

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?" +
			"Have you got access to is?\n")
		return
	}
	defer file.Close()

	// Gets a new reader
	reader := bufio.NewReader(file)

	// Initialize a new buffer
	block := make([]byte, 1)

	// Initialize a buffer of the read
	var buffer bytes.Buffer

	// Loop reads
	for {
		n, err := reader.Read(block)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Read File Error: %s\n", err)
		}
		if n == 0 || err == io.EOF {
			break
		}
		buffer.WriteString(string(block))
	}

	fmt.Printf("%s\n", buffer.String())
	err = ioutil.WriteFile(outputFile, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write File Error: %s\n", err)
	}
}
