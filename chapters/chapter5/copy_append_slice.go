package main

import (
	"fmt"
)

func copyAppend() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	//copy(dest,sour []T) int
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	var str1 = "Hello World111"
	var byteArr = make([]byte, 20, 20)
	//str1此处被转换成了byte[]类型
	p := copy(byteArr, str1)
	fmt.Printf("The string is %s!\n", str1)
	fmt.Println(string(byteArr))
	fmt.Println(p)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
