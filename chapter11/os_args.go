package main

import (
	"fmt"
	"os"
	"strings"
)

func osArgs() {
	str := ""
	fmt.Println(os.Args)
	for _, str := range os.Args {
		fmt.Println(str)
	}
	fmt.Println("The input string is :")
	if len(os.Args) > 1 {
		str += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(str)
}
