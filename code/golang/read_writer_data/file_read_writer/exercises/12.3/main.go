package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title    string
	Price    float64
	Quantity int
}

type books []Book

func main() {
	// Open the file
	file, err := os.Open("code/golang/read_writer_data/file_read_writer/exercises/12.3/books.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open the file error: %s", err)
	}

	// Close the open file
	defer file.Close()

	// Initialize sets
	bks := make(books, 0)

	// Gets a new reader
	reader := bufio.NewReader(file)

	// Loop read
	for {
		// Read line
		readString, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// Processing read data rows
		sections := strings.Split(readString, ";")

		if len(sections) == 3 {
			// Convert price data type
			price, err := strconv.ParseFloat(sections[1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "string parse to float64 error: %s", err)
				return
			}

			// Convert quantity data type
			quantity, err := strconv.Atoi(strings.TrimSpace(sections[2]))
			if err != nil {
				fmt.Fprintf(os.Stderr, "string parse to integer error: %s", err)
			}

			// Append to set
			bks = append(bks, Book{Title: sections[0], Price: price, Quantity: quantity})
		}
	}

	// Loop print
	fmt.Println("We have read the following books information from the file:")
	for _, v := range bks {
		fmt.Println(v)
	}
}
