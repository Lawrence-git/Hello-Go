package main

import (
	"fmt"
)

func forDemo() {
	// for i := 0; i < 6; i++ {
	// 	fmt.Printf("This is the %d!\n", i)
	// }
	forRange()
}

func forDemo() {
	var str1 = "Go is a beautiful language!"
	fmt.Printf("%s:length is %d \n", str1, len(str1))
	for i := 0; i < len(str1); i++ {
		fmt.Printf("The %d character is %c\n", i, str1[i])
	}

	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

//forDemo2等价于下面的forDemo3
func forDemo2() {
	i := 5
	for i > 0 {
		fmt.Printf("The variable i is now: %d\n", i)
		i--
	}
}

func forDemo3() {
	for i := 5; i > 0; {
		fmt.Printf("The variable i is now: %d\n", i)
		i--
	}
}

//for ix, val := range coll { }
// val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
func forRange() {
	str := "Hello World!"
	for pos, val := range str {
		fmt.Printf("%d :\t%c\n", pos, val)
	}
	for _, val := range str {
		fmt.Printf("%c\n", val)
	}
}
