package main

/**
当标识符以大写开头，则可以被外部包引用(类似与public)
当标识符以小写头，对包外不可见(类似与private)
*/
import "fmt"

// var flag bool = true
//var name type (不指定类型，使用默认值)
var flag bool

// var name = value (不指定类型,根据值自行判定变量类型)
var str = "Hello Go!"

//var vname1, vname2, vname3 type 同一个类型的多变量声明，
var flag1, flag2, flag3 bool

var str1, int1, bool1 = "String1", 100, true

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	struct1 = 300
	struct2 = "struct2"
)

//  :=  初始化声明,根据所给值的类型指定变量的type
func identifier() {
	vname1, vname2, vname3 := "String2", 200, true //这种不带声明的变量定义只能出现在函数体内
	fmt.Println("flag : ", flag)
	fmt.Println("str : ", str)
	fmt.Println("flag1:", flag1, "\tflag2:", flag2, "\tflag3:", flag3)
	fmt.Println("str1:", str1, "\tint1:", int1, "\tbool1:", bool1)
	fmt.Println(vname1, vname2, vname3)
	fmt.Println(struct1, struct2)
	fmt.Println(&struct1) // &i : 内存地址
}
