package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	var intP *int
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	intP = &i1
	fmt.Printf("The value of point intP is : %d\n", *intP)
	c := *intP + 1
	// d := *intP++  //在go中，这种用法是不合法的
	fmt.Println(c)
	//无法得到一个常量的地址
	// const i = 5
	// ptr := &i   //error: cannot take the address of i
	// ptr2 := &10 //error: cannot take the address of 10

	//对一个空指针的反向引用是不合法的，并且会使程序崩溃
	//runtime error: invalid memory address or nil pointer dereference
	var p *int
	*p = 0
}
