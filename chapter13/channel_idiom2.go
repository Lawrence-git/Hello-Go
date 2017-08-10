package main

import (
	"fmt"
	"time"
)

func channelIdiom2() {
	ch := pump()
	suck(ch)
	time.Sleep(1e6)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
			fmt.Println("input : ", i)
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
