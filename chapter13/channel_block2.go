package main

import (
	"fmt"
	"time"
)

func channelBlock2() {
	ch1 := make(chan int)
	go pumpChannel(ch1)
	go consumerChannel(ch1)
	time.Sleep(10 * 1e9)
}

func pumpChannel(ch chan int) {
	for i := 0; ; i++ {
		//发送操作,在接收者准备好之前是阻塞的
		fmt.Println("Send i = ", i)
		ch <- i
	}
}

func consumerChannel(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
