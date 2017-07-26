package main

import (
	"fmt"
)

//可以使用file这样小写的private的struct和大写的工厂方法，来强制用户使用工厂构造方法
type File struct {
	value    int
	fileName string
}

func NewFile(value int, fileName string) *File {
	if value < 0 {
		return nil
	}
	return &File{value, fileName}
}

func structFactory() {
	f := NewFile(10, "Hello World")
	fmt.Println(f)
}
