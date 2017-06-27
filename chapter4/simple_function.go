package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println(multilply3Nums(1, 3, 5))
	fmt.Println(getXandY(1, 2))
	fmt.Println(getXandY2(1, 2))
}

func multilply3Nums(a, b, c int) int {
	return a * b * c
}

/*
	相当于return x,y,定义了返回值名称时，return语句可以省写
	值的注意的是,此时return 1,2;返回的是1,2 而不是x,y
*/
func getXandY(input1, input2 int) (x int, y int) {
	y = input2 * 10
	x = input1 * 10
	return
}

//可以省略返回值的名字
func getXandY2(input1, input2 int) (int, int) {
	return input1 * 10, input2 * 10
}
