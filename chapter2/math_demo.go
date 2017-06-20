package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
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
