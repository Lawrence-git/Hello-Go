package main

import (
	"fmt"
	"strings"
)

func parameterDemo() {
	// result := callFunc(2, add)
	// fmt.Println("In main,result : ", result)

	//104,h	101,e	108,l	108,l	111,o
	str := "hello"
	f := func(c rune) bool {
		fmt.Printf("%d,%c\n", c, c)
		if c < 109 {
			return strings.ContainsRune("q", c)
		}
		return strings.ContainsRune("hello", c)
	}
	flag := strings.IndexFunc(str, f)
	fmt.Println(flag)
}

func add(a, b int) (sum int) {
	fmt.Println("Handler a + b")
	return a + b
}

func callFunc(x int, function func(int, int) (sum int)) (sum int) {
	fmt.Println("Before call function")
	sum = function(x, 2)
	fmt.Println("After call function")
	return
}
