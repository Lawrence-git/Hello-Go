package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	// A a
	A = iota
	//B b
	B
)

const (
	// E e
	E = iota
	// F f
	F
)

func mathDemo() {

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt16)
	var a = 65
	//b为string型65，c为字符串型A
	b := strconv.Itoa(a)
	c := string(a)
	fmt.Println("b :  " + b + ", c : " + c)
	var d, _ = strconv.Atoi(b)
	var e, _ = strconv.Atoi("65")
	fmt.Println(d)
	fmt.Println(e)
}
