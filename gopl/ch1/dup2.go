package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup2Main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	if input.Scan() {
		counts[input.Text()]++
	}
}
