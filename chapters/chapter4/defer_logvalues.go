package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 0, io.EOF
}

func deferLogDemo() {
	n, e := func1("Hello!")
	if e != nil {
		log.Println("Error:", e)
		return
	}
	log.Println("n:", n)
}
