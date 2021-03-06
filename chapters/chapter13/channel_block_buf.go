package main

import (
	"fmt"
	"time"
)

//buf(async):1038241  //sync:684632
func channelBlockBuf() {
	ch := make(chan string)
	go sendDataBuf(ch)
	go getDataBuf(ch)
	time.Sleep(2 * 1e9)
}

func sendDataBuf(ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprint("Hello World ", i)
	}
}

func getDataBuf(ch chan string) {
	for {
		str := <-ch
		fmt.Println(str)
	}
}
