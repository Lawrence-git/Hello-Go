package main

import (
	"fmt"
	"time"
)

func goroutine2() {
	//定义channel
	// var ch1 chan string
	// ch1 = make(chan string)
	ch := make(chan string)
	go getData(ch)
	go sendData(ch)
	time.Sleep(1 * 1e9)
}

func sendData(ch chan string) {
	ch <- "Hello World"
	ch <- "Beijing"
	ch <- "Tokyo"
	ch <- "WashingTon"
}

func getData(ch chan string) {
	var chanString string
	for {
		chanString = <-ch
		fmt.Printf("%s\n", chanString)
	}
}
