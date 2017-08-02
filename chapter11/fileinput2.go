package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func fileinput2() {
	inputFile := "demo1.txt"
	outputFile := "demo2.txt"
	buf, err1 := ioutil.ReadFile(inputFile)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err1)
		return
	}
	fmt.Printf("%s\n", string(buf))
	err2 := ioutil.WriteFile(outputFile, buf, 0644)
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Write File Filed: %s\n", err1)
	}
}
