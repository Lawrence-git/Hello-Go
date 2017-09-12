package main

import (
	"fmt"
)

func bubbleSortDemo() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)

	fmt.Println("Bubble Sort")

	length := len(a)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
