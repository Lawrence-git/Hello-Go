package main

import (
	"fmt"
	"reflect"
)

func sliceDemo2() {
	//切片是一个数组片段的描述。它包含了指向数组的指针，片段的长度， 和容量（片段的最大长度）
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(reflect.TypeOf(arr1)) //[3]int
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice2)) //[]int

	slice3 := arr1[:cap(arr1)]
	fmt.Println(slice3)
	fmt.Println(reflect.TypeOf(slice3)) //[]int
}
