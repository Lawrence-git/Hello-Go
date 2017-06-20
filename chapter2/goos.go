package main

import (
	"fmt"
	"os"
)

var (
	//HOME home路径
	HOME = os.Getenv("HOME")
	//USER current user
	USER = os.Getenv("USER")
	//GOROOT go的root路径
	GOROOT = os.Getenv("GOROOT")
)

func goos() {
	var user string = os.Getenv("USER")
	fmt.Printf("The current user is: %s!\n", user)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
