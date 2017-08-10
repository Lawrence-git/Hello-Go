package main

import "fmt"

func channelSemaphore() {
	ch := make(chan string)
	go doPrint1(ch)
	go doPrint1(ch)
}

func doPrint1(ch chan string) {
	ch <- "print1"
	fmt.Println("Hello World,Print 1")
}

func doPrint2(ch chan string) {
	ch <- "print2"
	fmt.Println("Hello World,Print 2")
}
