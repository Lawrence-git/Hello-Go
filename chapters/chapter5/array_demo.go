package main

import (
	"fmt"
)

func arrayDemo() {
	var arr1 [5]int
	fmt.Println(arr1)

	a := [5]int{1, 2}
	fmt.Println(a)

	b := [20]int{1, 2, 3: 1}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)

	p := &b
	p[9] = 10
	fmt.Println("p:", *p)
	fmt.Println("b:", b)

	d := [...][3]int{
		{1, 2, 3},
		{2, 3, 4}}
	fmt.Println("d:", d)
	//创建数组切片，并仅初始化其中的部分元素，数组切片的len将根据初始化的元素确定
	var array6 = []string{4: "Smith", 2: "Alice"}
	fmt.Printf("array6--- type:%T \n", array6)
	rangeObjPrint(array6)
}

//输出字符串数组切片
func rangeObjPrint(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%s\n", i, v)
	}
}
