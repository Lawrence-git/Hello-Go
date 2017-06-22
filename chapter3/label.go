package main

import "fmt"

func main() {
	gotoDemo2()
	forDemoL1()
	forDemoL2()
}

//go的标签用户和java一致
func labelDemo1() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

//不推荐使用goto和标签，因为代码的可读性不好
func gotoDemo() {
	i := 0
HERE:
	i++
	if i == 5 {
		println("i == 5")
		return
	}
	println("i == ", i)
	goto HERE
}

func gotoDemo2() {
	var b int
	i := 10
	goto LABEL2
	b = 9 //unreachable code!!
LABEL2:
	println("i == ", i)
	println("b == ", b)
}

func forDemoL1() {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}

func forDemoL2() {
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
