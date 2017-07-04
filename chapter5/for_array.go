package main

import (
	"fmt"
)

func forArray() {
	var arr1 [5]int

	fmt.Println("Init Array")
	for i := 0; i < len(arr1); i++ {
		arr1[i] = (i + 1) * 10
	}
	fmt.Println(arr1)

	fmt.Println("Print Array")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[i] = %d \t", arr1[i])
	}
	fmt.Println()

	fmt.Println("For Range Print Array")
	//range是一个内置函数，返回两个值 index和value
	//可以省略掉返回的value,写作for i : range arr
	//如果不需要index，则需要写作 for _,v : range arr
	for i, v := range arr1 {
		fmt.Printf("arr1[%d] = %d \t", i, v)
	}
	fmt.Println()

	a := [...]string{"a", "b", "c", "d"}
	for _, v := range a {
		fmt.Println("Array item", 0, "is", v)
	}
}
