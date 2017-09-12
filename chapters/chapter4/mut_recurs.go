package main

import (
	"fmt"
)

func mutRecurs() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is even: is %t\n", 17, even(17)) // 17 is even: is false
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(cutDown(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(cutDown(nr) - 1)
}

func cutDown(nr int) int {
	if nr <= 0 {
		return 0
	}
	return nr
}
