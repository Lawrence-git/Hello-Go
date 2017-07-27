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

type TT float64

func (t TT) String() string {
	return fmt.Sprintf("%v", t)
}

func structsString() {
	s := Student{"2012551435", "xia"}
	//因为是实例，所以没有调用String()方法
	fmt.Println(s)
	//因为是指针，实际上调用了String()方法
	fmt.Println(&s)
	//显示调用
	fmt.Println(s.String())
	fmt.Sprintf()

	t := TT{3.32}
	fmt.Println(t)
}
