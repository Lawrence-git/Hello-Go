package main

import (
	"fmt"
	"time"
)

func goroutineSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pumpData1(ch1)
	go pumpData1(ch2)
	go suckData(ch1, ch2)
	time.Sleep(1 * 1e9)
	time.Tick(1e9)
}

func pumpData1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 3
	}
}

func suckData(ch1, ch2 chan int) {
	for {
		select {
		case u := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", u)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
