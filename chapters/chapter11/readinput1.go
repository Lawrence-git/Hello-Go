package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func readInput1() {
	fmt.Println("Please enter your firstName:")
	fmt.Scanln(&firstName)
	fmt.Println("Please enter your lastName")
	fmt.Scanln(&lastName)
	fmt.Printf("Your name is %s %s\n", firstName, lastName)
	fmt.Println("start format input")
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
