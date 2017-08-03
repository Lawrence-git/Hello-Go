package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, _ := os.Open("products.txt")
	outputFile, _ := os.OpenFile("productsC.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError != nil {
			fmt.Println("Error", readerError)
			return
		}
		outputString := string(inputString) + "\n"
		fmt.Printf(outputString)
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputWriter.Flush()
	}
	fmt.Println("Conversion done")
}
