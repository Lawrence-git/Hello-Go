package main

import (
	"fmt"
)

type Student struct {
	studentSno string
	name       string
}

func (stu *Student) String() string {
	return "The student " + stu.name + " sno is " + stu.studentSno + "."
}

func main() {
	s := Student{"2012551435", "xia"}
	//因为是实例，所以没有调用String()方法
	fmt.Println(s)
	//因为是指针，实际上调用了String()方法
	fmt.Println(&s)
	//显示调用
	fmt.Println(s.String())
}
