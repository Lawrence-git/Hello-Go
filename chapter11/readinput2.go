package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufferReader *bufio.Reader

func readInput2() {
	bufferReader = bufio.NewReader(os.Stdin)
	str, err := bufferReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", str)
	} else {
		fmt.Println(err)
	}
}
