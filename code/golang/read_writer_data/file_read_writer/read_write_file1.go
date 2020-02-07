package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "code/golang/read_writer_data/file_read_writer/products.txt"
	outputFile := "code/golang/read_writer_data/file_read_writer/products_copy.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", err)
	}

	fmt.Printf("%s\n", string(bytes))
	err = ioutil.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write File Error: %s\n", err)
	}
}
