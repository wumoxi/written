package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(f *os.File) {
	const NBUF = 512
	var buff [NBUF]byte
	for {
		switch nr, err := f.Read(buff[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buff[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err == io.EOF {
			continue
		}
		cat(file)
		file.Close()
	}
}
