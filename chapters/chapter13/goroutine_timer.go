package main

import (
	"fmt"
	"time"
)

func main() {
	//以 duration 为周期给返回的通道发送时间
	tick := time.Tick(1e9)
	//After() 只发送一次
	boom := time.After(5e9)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
