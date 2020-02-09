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
		buff, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *isShowLineNumber {
			number++
			fmt.Fprintf(os.Stdout, "%5d %s", number, buff)
		} else {
			fmt.Fprintf(os.Stdout, "%s", buff)
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
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error: %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		cat(bufio.NewReader(file))
		file.Close()
	}
}
