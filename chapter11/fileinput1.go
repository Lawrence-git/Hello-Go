package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileinput1() {
	file, err := os.Open("demo1.txt")
	if err != nil {
		fmt.Println("Something error is happened!")
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	inputReader := bufio.NewReader(file)
	buf := make([]byte, 256)
	var str string = ""
	for {
		n, _ := inputReader.Read(buf)
		str = str + string(buf)
		if n == 0 {
			break
		}
	}
	fmt.Printf(str)
}
