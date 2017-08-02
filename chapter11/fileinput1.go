package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func fileinput1() {
	file, err := os.Open("readinput11.go")
	if err != nil {
		fmt.Println("Something error is happened!")
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	inputReader := bufio.NewReader(file)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError != nil {
			if readerError == io.EOF {
				return
			} else {
				fmt.Println("Something error is happened when reading file!")
			}
		}
		fmt.Println(inputString)
	}
}
