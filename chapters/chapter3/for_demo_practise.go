package main

import (
	"fmt"
)

func forDemoPractise() {
	demo2()
}

//0 0 0 0 0
func demo1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	fmt.Println()
}

//0 1 2 3 4 5 6 ......
func demo2() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}
