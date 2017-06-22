package main

import (
	"fmt"
)

func switchDemo() {
	switchInt(1)
	switchInt(2)
	switchInt(3)
	switchInt(4)
	switchInt(5)
	switchInt(6)
}

func switchInt(aInt int) {
	switch aInt {
	case 0:
		fmt.Println("000000")
	case 1:
		fmt.Println("1111111")
	case 2:
	case 3:
		fmt.Println("33333333")
	case 4, 5:
		fmt.Println("444455555")
	default:
		fmt.Println("Default")
	}
}
