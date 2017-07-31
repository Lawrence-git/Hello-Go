package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (square *Square) Area() float32 {
	return square.side * square.side
}

type Rectangle struct {
	length, width float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.length * rectangle.width
}

type Circle struct {
	radius float32
}

func (circle *Circle) Area() float32 {
	return math.Pi * circle.radius * circle.radius
}

func simpleDemo() {
	//接口类型的指针，最终指向了实现类
	var areaIntf Shaper
	areaIntf = &Square{10.01}
	fmt.Println(areaIntf.Area())

	areaIntf = &Rectangle{10.01, 10.02}
	fmt.Println(areaIntf.Area())

	if t, ok := areaIntf.(*Rectangle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Rectangle, *Circle:
		fmt.Printf("Type Rectangle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("nil value: nothing to check?\n")
	}
}
