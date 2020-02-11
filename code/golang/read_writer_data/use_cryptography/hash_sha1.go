package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	// Gets a new hash
	hasher := sha1.New()

	// Initialize byte slice
	var b []byte

	// Hash write string. (散列写字符串)
	_, err := io.WriteString(hasher, "test")
	if err != nil {
		log.Fatalf("write string error: %s\n", err)
	}

	// Sum appends the current hash to b and returns the resulting slice. (Sum将当前散列附加到b并返回结果片)
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	// Hash write byte slice. (散列写字节片)
	hasher.Reset()
	data := []byte("We shall overcome!")
	_, err = hasher.Write(data)
	if err != nil {
		log.Fatalf("write byte slice error: %s\n", err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
