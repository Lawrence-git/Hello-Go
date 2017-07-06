package main

import (
	"fmt"
)

func main() {
	var str = "Hello World"
	str1, str2 := splitString(str, 4)
	fmt.Println(str1)
	fmt.Println(str2)
}

func splitString(str string, index int) (str1, str2 string) {
	charSlice := []byte(str)
	str1 = string(charSlice[:index])
	str2 = string(charSlice[index:])
	return
}
