package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world - unbuffered")

	// buffered: os.Stdout implements io.Writer
	buff := bufio.NewWriter(os.Stdout)

	// and now so does buff.
	fmt.Fprintf(buff, "%s\n", "hello world - buffered")

	// Flush buffer data to buff
	buff.Flush()
}
