package main

import "fmt"

type Sorter interface {
	Len() int
	Swap(a, b int)
	Less(a, b int) bool
}

func sort(data Sorter) {
	for pass := 0; pass < data.Len(); pass++ {
		for i := pass; i < data.Len(); i++ {
			if data.Less(i, pass) {
				data.Swap(i, pass)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type Array []int

func (array Array) Len() int {
	return len(array)
}

func (array Array) Less(a, b int) bool {
	return array[a] < array[b]
}

func (array Array) Swap(a, b int) {
	array[a], array[b] = array[b], array[a]
}

func SortDemo() {
	var array Array
	array = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	fmt.Println(array)
	sort(array)
	fmt.Println(array)
}
