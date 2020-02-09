package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var isShowLineNumber = flag.Bool("n", false, "show file content of line number")

func cat(r *bufio.Reader) {
	var number int64
	for {
		bytes, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *isShowLineNumber {
			number++
			fmt.Fprintf(os.Stdout, "%5d %s", number, bytes)
		} else {
			fmt.Fprintf(os.Stdout, "%s", bytes)
		}
	}
	return
}

func main() {
	flag.PrintDefaults()
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: error: %s\n", os.Args[0], flag.Arg(i), err)
			return
		}
		cat(bufio.NewReader(file))
	}
}
