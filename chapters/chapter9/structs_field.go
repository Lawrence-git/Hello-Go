package main

import (
	"fmt"
)

type human struct {
	name string
	age  int
}

func structsField() {
	xiaoming := new(human)
	xiaoming.name = "xiaoming"
	xiaoming.age = 20
	fmt.Printf("The name is %s. The age is %d.\n", xiaoming.name, xiaoming.age)
	//更简短的初始化结构体变量
	xiaohong := human{"xiaohong", 21}
	fmt.Printf("The name is %s. The age is %d.\n", xiaohong.name, xiaohong.age)

	// xiaojun := person{"xiaojun", 21}
	// xiaojun := person{name: "xiaojun"}
	xiaojun := human{age: 21}
	fmt.Printf("The name is %s. The age is %d.\n", xiaojun.name, xiaojun.age)
}
