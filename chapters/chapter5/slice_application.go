package main

import (
	"fmt"
	"sort"
	"strings"
)

func sliceApplicationDemo() {
	var str = "Hello World"
	//丛字符串生成切片
	// c := []byte(str)
	// fmt.Printf("%d:%c\n", 10, c[10])
	// fmt.Println("================================")
	// forRangeString(str)

	//获取字符串的子串
	// substr := str[2:]
	// fmt.Println(substr)

	//go语言中字符串是不可变的。只有先将其变为数组，修改其中的值，最后再将其转换回字符串
	// charSlice := []byte(str)
	// charSlice[0] = 'C'
	// str = string(charSlice)
	// fmt.Println(str)

	//str < "Hello Test"
	n := strings.Compare(str, "Hello Test")
	fmt.Println(n)

	intArr := []int{1, 3, 7, 5, 6}
	fmt.Println(intArr)
	flag1 := sort.IntsAreSorted(intArr)
	fmt.Println("Is sorted?", flag1)

	sort.Ints(intArr)

	fmt.Println(intArr)
	flag2 := sort.IntsAreSorted(intArr)
	fmt.Println("Is sorted?", flag2)
}

func forRangeString(str string) {
	for i, c := range str {
		fmt.Printf("%d:%c\t", i, c)
	}
	fmt.Println()
}
