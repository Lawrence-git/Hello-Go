package main

import (
	"fmt"
)

type innerS struct {
	in1 int
	in2 int
}

//go语言更青睐组合
type outer struct {
	a int
	b int
	innerS
}

func structsAnonymousField() {
	fmt.Println()
	innerS1 := innerS{1, 2}
	outer1 := outer{1, 2, innerS1}
	outer1.in1 = 10
	fmt.Println(outer1)
}
