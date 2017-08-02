package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.OpenFile("demo2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello World\n")
	}
	writer.WriteString("Hello World")
	writer.Flush()
}
