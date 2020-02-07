package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "code/golang/read_writer_data/file_read_writer/products2.txt"
	fmt.Printf("file path of the base: %s\n", filepath.Base(path))
}
