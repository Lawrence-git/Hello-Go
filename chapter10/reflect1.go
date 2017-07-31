package main

import (
	"fmt"
	"reflect"
)

func reflect1() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) //type:float64
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)                      //value:3.4
	fmt.Println("type:", v.Type())                //type:float64
	fmt.Println("kind:", v.Kind())                //kind:float64
	fmt.Println("value:", v.Float())              //value:3.4
	fmt.Println(v.Interface())                    //3.4
	fmt.Printf("value is %5.2e\n", v.Interface()) //value is 3.40e+00
	y := v.Interface().(float64)
	fmt.Println(y) //3.4
}
