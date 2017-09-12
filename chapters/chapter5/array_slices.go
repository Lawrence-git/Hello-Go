package main

import "fmt"

func arraySliceDemo2() {
	var arr1 [6]int
	var slice1 = arr1[3:5] // index 5 niet meegerekend!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice1 at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	//cap(s) 就是从 s[0](切片的start位置!!!) 到数组末尾的数组长度
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice:
	slice2 := slice1[0:3]
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice2 at %d is %d\n", i, slice2[i])
	}
	fmt.Printf("The length of slice2 is %d\n", len(slice2))
	fmt.Printf("The capacity of slice2 is %d\n", cap(slice2))

	slice3 := arr1[:2]
	fmt.Printf("The length of slice3 is %d\n", len(slice3))   //2
	fmt.Printf("The capacity of slice3 is %d\n", cap(slice3)) //len(arr)

	slice4 := arr1[2:]
	fmt.Printf("The length of slice4 is %d\n", len(slice4))   //len(arr)-2
	fmt.Printf("The capacity of slice4 is %d\n", cap(slice4)) //len(arr)-2
}
