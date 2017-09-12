package main

import (
	"fmt"
	"runtime"
)

/*
  if-else 的常见写法，通过process()函数获取value的初始值，在与某个值比较
  if value := process(data); value > max {
    ...
  if value := process(data); value < min {
    ...
  }
*/

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		//Sprintf根据format参数生成格式化的字符串并返回该字符串。
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func ifDemo2() {
	fmt.Println(prompt)
	fmt.Println(Abs(-100))
	fmt.Println(isGreaterThan10(10))

	// var cond int = 0
	//此处 := 初始化的cond是一个只存在于 if-else 分支内的局部变量
	if cond := 15; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
	// fmt.Printf("cond is %d\n", cond)
}

// Abs 获取x的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreaterThan10(x int) bool {
	/*
	   *!if 可以包含一个初始化语句
	   if initialization; condition {
	       // do something
	   }
	*/
	if val := 10; x > val {
		return true
	}
	return false
}
