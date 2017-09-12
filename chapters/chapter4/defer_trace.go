package main

import "fmt"

func deferTrace() {
	b()
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}
