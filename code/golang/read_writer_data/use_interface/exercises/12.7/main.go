package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("code/golang/read_writer_data/use_interface/exercises/12.7/goprogram.txt")
	outputFile, _ := os.OpenFile("code/golang/read_writer_data/use_interface/exercises/12.7/goprogramT.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		fmt.Println(inputString)
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Flush the buffer data to a file
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")
}
