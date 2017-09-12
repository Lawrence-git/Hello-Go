package main

import (
	"fmt"
	"os"
)

func init() {
	user := os.Getenv("USER")
	fmt.Println(user)
	//panic不会影响defer语句的执行
	defer fmt.Println("Exiting init")
	if user == "RuzzZZ" {
		panic("error:current user is " + user)
	}
}

func panicDemo() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
