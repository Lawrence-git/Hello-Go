package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b B) changeValue() { b.thing = 1 }

//changePointer接受的是一个指针!!!!
func (b *B) changePointer() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B
	b1.changeValue()
	fmt.Println(b1.write())

	b1.changePointer()
	fmt.Println(b1.write())
}
