package main

import (
	"fmt"
	"reflect"
)

func makeAlice() {
	//func make([]T, len, cap)，其中 cap 是可选参数。
	var slice1 = make([]int, 10, 20)
	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = i * 10
	}
	fmt.Println(slice1)
	fmt.Println(reflect.TypeOf(slice1))

	//func new(type)
	slice2 := new([1]int)[:]
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice2))

	fmt.Println("======================================================================")
	// var p *[]int = new([]int) // *p == nil; with len and cap 0
	p := new([1]int)
	for i := range *p {
		p[i] = i + 1
	}
	fmt.Println(*p)
}
