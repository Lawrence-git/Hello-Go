package main

//go语言中,方法 != 函数(!!!!)
//go语言中，方法是包含了接收者参数的函数
import (
	"fmt"
)

type twoInts struct {
	a int
	b int
}

func structs_method() {
	ti := twoInts{10, 20}
	fmt.Println(ti.addA(10))
	fmt.Println(ti.addAB(10))
}

//接受值
func (ti twoInts) addA(value int) (sum int) {
	return ti.a + value
}

func (ti twoInts) addAB(value int) (sum int) {
	return ti.a + value + ti.b
}
