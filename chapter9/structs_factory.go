package main

import (
	"fmt"
)

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

func main() {
	f := NewFile(10, "Hello World")
	fmt.Println(f)
}
