package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

//子类的Abs()方法,如果要调用父类的Abs()方法， son.Father.Method()
func (p *NamedPoint) Abs() float64 {
	return math.Sqrt(p.x*p.x+p.y*p.y) * 100
}

func structsOverride() {
	np := NamedPoint{Point{3, 4}, "HelloWorld"}
	fmt.Println(np.Abs())
	fmt.Println(np.Point.Abs())
}
