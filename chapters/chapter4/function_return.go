package main

import (
	"fmt"
	"log"
	"runtime"
)

func returnDemo() {
	where()
	f1 := add2()
	log.SetFlags(log.Llongfile)
	log.Println("")
	f2 := addA(2)
	fmt.Println(f1(3))
	var logWhere = log.Println
	logWhere()
	fmt.Println(f2(4))
	where()
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func addA(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d\n", file, line)
}
