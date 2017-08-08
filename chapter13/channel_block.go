package main

import (
	"fmt"
	"time"
)

func channelBlock() {
	ch1 := make(chan int)
	go pump(ch1)
	time.Sleep(5 * 1e9)
	//接收操作是阻塞的,直到发送者可用
	fmt.Println(<-ch1)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		//发送操作,在接收者准备好之前是阻塞的
		fmt.Println("Send i = ", i)
		ch <- i
	}
}
