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
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("No args was input")
	}
	for _, v := range flag.Args() {
		file, err := os.Open(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], v, err.Error())
			continue
		}
		cat(bufio.NewReader(file))
	}

}
