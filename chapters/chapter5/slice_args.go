package main

import (
	"fmt"
)

func sliceArgs() {
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)
	//将切片作为参数调用函数sum,
	//因为切片传递的是函数的指针，所以arr数组的值回被改变
	sum(arr[:])
	fmt.Println(arr)
}

func sum(a []int) {
	a[2] = 100
}
