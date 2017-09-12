package main

import (
	"fmt"
)

func main() {
	// fallthroughDemo()
	switchInt(3)
}

//go语言接受任意类型的switch变量
func switchInt(aInt int) {
	switch aInt {
	case 0:
		fmt.Println("000000")
	case 1:
		fmt.Println("1111111")
	case 2:
	case 3:
		fmt.Println("33333333")
	case 4, 5:
		fmt.Println("444455555")
	default:
		fmt.Println("Default")
	}
}

//当任一分支的测试结果为 true 时，该分支的代码会被执行
//实际上就是 switch-true 的过程
func switchBool(aInt int) {
	switch {
	case aInt > 0:
		fmt.Println("aINt > 0")
	case aInt == 0:
		fallthrough
	default:
		fmt.Println("aInt < 0")
	}
}

//与上面的 switchBool() 并无本质区别
//只是先定义了一个switch的局部变量，注意末尾的 ";"号
func switchBool2(aInt int) {
	switch result := always10(); {
	case result > 0:
		fmt.Println("aINt > 0")
	case result == 0:
		fmt.Println("aInt == 0")
	default:
		fmt.Println("aInt < 0")
	}
}

//可以使用fallthrough强制执行后面的case代码(不经过判断(case)部分！！)
func fallthroughDemo() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func always10() (aInt int) {
	return 10
}
