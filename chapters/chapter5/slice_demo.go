package main

import (
	"fmt"
)

func sliceDemo() {
	arr1 := [...]int{1, 5, 10, 15, 20}
	// arr[start:end] : 从start为到end-1的索引
	//start为0可以省略start，end为len(arr)可以省略end
	var slice1 = arr1[:] //等价于var slice1 = arr1[0:len(arr1)]
	for _, v := range slice1 {
		fmt.Println(v)
	}
	fmt.Printf("capacity : %d; length : %d\n", cap(slice1), len(slice1))
	fmt.Println("==============================")

	var slice2 []int
	slice2 = arr1[:len(arr1)-1] //等价于slice2 = arr1[0 : len(arr1)-1]
	for _, v := range slice2 {
		fmt.Println(v)
	}
	fmt.Printf("capacity : %d; length : %d\n", cap(slice2), len(slice2))
	fmt.Println("==============================")

	slice3 := arr1[0 : len(arr1)-2]
	for _, v := range slice3 {
		fmt.Println(v)
	}
	fmt.Printf("capacity : %d; length : %d\n", cap(slice3), len(slice3))
}
