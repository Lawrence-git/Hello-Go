package main

import (
	"fmt"
	"unsafe"
)

//const identifier [type] = value (类型type可以省略)
const booleanTrue bool = true
const booleanFalse = false

//常量还可以用作表示枚举
const (
	UNKNOW = 0
	FEMALE = 1
	MALE   = 2
)

const (
	a = iota
	b = iota
	c = iota
)

func constVariable() {
	fmt.Println(booleanTrue, booleanFalse)
	fmt.Println(UNKNOW, FEMALE, MALE)
	const str = "Ssdasdasw"
	fmt.Println(len(str), "\t", unsafe.Sizeof(str))
	const flag = 300
	fmt.Println(max(100, flag))
	//iota
	fmt.Println(a, b, c)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
