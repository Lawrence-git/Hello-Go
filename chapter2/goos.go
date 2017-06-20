package main

import (
	"fmt"
	"os"
)

func main() {
	var user string = os.Getenv("USER")
	fmt.Printf("The current user is: %s!\n", user)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
