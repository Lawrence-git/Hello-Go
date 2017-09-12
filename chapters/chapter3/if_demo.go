package main

import (
	"fmt"
	"runtime"
)

func ifDemo() {
	var result = testIf()
	fmt.Println(result)

	if runtime.GOOS == "window" {
		fmt.Println("Windows!")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("darwin!")
	} else {
		fmt.Println("Other:", runtime.GOOS)
	}
}

func testIf() bool {
	bool1 := true
	if !bool1 {
		fmt.Println("True")
		return true
	}
	fmt.Println("False")
	return false
}
