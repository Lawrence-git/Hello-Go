package main

import "fmt"

func badCall() {
	fmt.Printf("Starting a bad call!\n")
	panic("Bad Call!")
}

func test() {
	defer func() {
		//recover()函数只能在defer修饰的函数中使用
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
