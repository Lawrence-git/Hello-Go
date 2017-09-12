package main

import "fmt"

func channelResult() {
	array := [...]int{1, 2, 3, 4}
	ch := make(chan int)
	go sum(ch, array)
	//do something else

	//因为没有缓冲区,所以goroutine是阻塞的
	sum := <-ch
	fmt.Println("sum is ", sum)
}

func sum(ch chan int, array [4]int) {
	sum := 0
	for _, v := range array {
		sum += v
	}
	ch <- sum
}
