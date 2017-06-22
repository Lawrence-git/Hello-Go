package main

import (
	"fmt"
)

func breakContinue() {
	continueDemo()
}

func breakDemo() {
	i := 10
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i <= 0 {
			break
		}
	}
}

func continueDemo() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Find a five!")
			continue
		}
		fmt.Printf("The num is %d\n", i)
	}
}
