package main

import (
	"fmt"
)

func multiRange() {
	// testMultiForRange()
	forRange()
}

func forRange() {
	items := [...]int{10, 20, 30, 40, 50}
	//value拿到的是值的copy，所以不能修改数组中的值
	for _, value := range items {
		value *= 2
	}
	fmt.Println(items)
	//数组内的值扩大两倍
	for index := range items {
		items[index] *= 2
	}
	fmt.Println(items)
}

func testMultiForRange() {
	arr := [3][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	fmt.Println(arr)
	fmt.Println("Multi for-range to print arr!")
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Println()
	}
}
