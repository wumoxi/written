package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buff, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buff)
	}
	return
}

var LineNumber = flag.Bool("n", false, "show file content of line number")

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error: %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(bufio.NewReader(file))
		file.Close()
	}
}
