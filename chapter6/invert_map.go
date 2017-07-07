package main

import (
	"fmt"
	"reflect"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	fmt.Println(reflect.TypeOf(invMap))
	for k, v := range barVal {
		if _, isPresent := invMap[v]; isPresent {
			fmt.Printf("The key %d is exist.The value %s can't be inserted!\n", v, k)
		} else {
			invMap[v] = k
		}
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v \n ", k, v)
	}
}
