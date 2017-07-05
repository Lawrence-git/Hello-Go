package main

import (
	"fmt"
)

func main() {
	// reslice(
	reslice2()
}

func reslice2() {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

func reslice() {
	arr := [10]int{1: 1, 6: 20, 3: 19, 7: 8}
	fmt.Println(arr)

	slice1 := arr[1:3]
	fmt.Println(slice1)
	//切片重组
	slice1 = slice1[0 : len(slice1)+2]
	fmt.Println(slice1)
}
