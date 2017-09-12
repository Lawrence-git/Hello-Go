package main

import "fmt"

func channelGoroutine3() {
	ch := make(chan string)
	//sendData1是一个协程
	go sendData1(ch)
	//getData1和main()在同一个线程内
	getData1(ch)
}

func sendData1(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData1(ch chan string) {
	for {
		input, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("%s\n", input)
	}
}
