package main

import (
	"fmt"
)

//go语言中数组是一种值类型,所以数组不能直接作为参数传入函数
//如果想在函数中对数组进行修改，可以将数组的地址作为参数传入
func main() {
	arr1 := [...]int{0, 5, 10, 15}
	fmt.Println(arr1)
	arr2 := arr1
	arr2[1] = 10
	fmt.Println(arr1[1])

	var arrKeyValue = []string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
	println("========================================================================================")
	arrayPoint1(&arr1)
	println("========================================================================================")
	arrayPoint2(&arr1)
}

func arrayPoint1(a *[4]int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t", a[i])
	}
	fmt.Println()
}

func arrayPoint2(a *[4]int) {
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
