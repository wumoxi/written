package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file(打开文件)
	file, err := os.Open("code/golang/read_writer_data/file_read_writer/products2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open the file error: %s\n", err)
		return
	}

	// Close the open file(关闭文件)
	defer file.Close()

	// Initializes the column read result(初始化列读取结果)
	var col1, col2, col3 []string

	// Loop read(循环读取)
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	// Print the column read result(打印列读取结果)
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
