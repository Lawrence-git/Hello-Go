package main

import "fmt"

func innerFuncDemo() {
	//new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）
	aIntPointer := new(int)
	fmt.Println(aIntPointer)
	fmt.Println(*aIntPointer)

	bInt := 1
	bIntPointer := &bInt
	fmt.Println(bIntPointer)
	fmt.Println(*bIntPointer)
}
