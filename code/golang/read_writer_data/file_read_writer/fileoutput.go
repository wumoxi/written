package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string

	// Open the file(打开文件)
	outputFile, outputError := os.OpenFile("code/golang/read_writer_data/file_read_writer/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}

	// Close the open file(关闭文件)
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello world\n"

	// Write multiple times through the loop to the write buffer (通过循环写入多次到写入缓冲)
	for i := 0; i < 10; i++ {
		times := strconv.Itoa(i)
		outputWriter.WriteString(times + ". " + outputString)
	}

	// Flush buffer data to a file(将缓冲区数据冲洗到文件)
	outputWriter.Flush()
}
