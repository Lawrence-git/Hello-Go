package main

import (
	"fmt"
	"reflect"
)

func reflectDemo() {
	var f float64 = 3.23
	fmt.Println(reflect.TypeOf(f))         //float64
	fmt.Println(reflect.ValueOf(f))        //3.23
	fmt.Println(reflect.ValueOf(f).Type()) //float64
	//Kind()方法:返回类型
	fmt.Println(reflect.TypeOf(f).Kind())
	fmt.Println(reflect.ValueOf(f).Kind())
}
