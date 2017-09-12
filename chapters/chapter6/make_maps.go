//make是一个引用类型，申明方式如下；在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
//	var map1 map[keytype]valuetype
// var map1 map[string]int
package main

import (
	"fmt"
)

func makeMaps() {
	// var mapLit map[string]int
	// mapLit = map[string]int{"one": 1, "two": 2, "three": 3}
	// mapLit["four"] = 4
	// fmt.Println(mapLit)

	//map的容量是动态伸缩的，但是可以指定初始化时的容量大小(类似java的arraylist)
	mapSlice := make(map[int]string, 100)
	mapSlice[1] = "one"
	fmt.Println(mapSlice)

	// mapExample := map[string]int{"one": 1, "two": 2}
	// fmt.Println(mapExample["one"])

	//func作为map的value，可以用来做分支结构
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	//value,isPresnt := map1[key1]
	value1, isPresent1 := mf[1]
	if isPresent1 {
		fmt.Println(value1())
	}
	if value, isPresent := mf[4]; isPresent {
		fmt.Println(value())
	}
	//删除元素，如果元素不存在，不会报错!!
	delete(mf, 1)
	fmt.Println(mf)
	//for range,注意：map是无序的!!!
	for key, value := range mf {
		fmt.Print("The key : ", key)
		fmt.Println("\tThe value of func(): ", value())
	}

	fmt.Println("======================================================")
	// map类型的数组切片
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
	for i := range items {
		for j := range items[i] {
			fmt.Println(items[i][j])
		}
	}

}
