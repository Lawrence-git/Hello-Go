package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader

func readInput3() {
	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Your name is %s", str)

	switch str {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Chris\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
