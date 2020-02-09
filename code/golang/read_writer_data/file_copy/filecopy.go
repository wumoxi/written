package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("code/golang/read_writer_data/file_copy/target.txt", "code/golang/read_writer_data/file_copy/source.txt")
	fmt.Println("Copy done!")
}

// CopyFile is the a file copy function
func CopyFile(dstName string, srcName string) (written int64, err error) {
	// Open the file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	// Close the open file
	defer src.Close()

	// Create a new file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	// Close the newly created file
	defer dst.Close()

	// Copy file
	return io.Copy(dst, src)
}
