package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readTxtDemo() {
	file, error := os.Open("products.txt")
	if error != nil {
		panic(error)
	}
	reader := bufio.NewReader(file)
	var title, price, quantity []string
	for {
		buf, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		temp := strings.Split(string(buf), ";")
		title = append(title, temp[0])
		price = append(price, temp[1])
		quantity = append(quantity, temp[2])
	}
	fmt.Println(title)
	fmt.Println(price)
	fmt.Println(quantity)
}
