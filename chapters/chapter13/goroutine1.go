package main

import (
	"fmt"
	"time"
)

func goroutine1() {

	fmt.Println("Start in the main!")
	//routine
	go longWait()
	go shaortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) //5 seconds
	fmt.Println("End of longWait()")
}

func shaortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // 2 seconds
	fmt.Println("End of shortWait()")
}
