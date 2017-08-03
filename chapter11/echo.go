package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var s string
	for i, v := range flag.Args() {
		if i > 0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}
		s += v
	}
	os.Stdout.WriteString(s)
}
