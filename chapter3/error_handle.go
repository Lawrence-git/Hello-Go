package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func errorHandler() {
	if aInt, err1 := strconv.Atoi("qwq"); err1 != nil {
		fmt.Printf("An error occured!%s\n", err1.Error())
	} else {
		fmt.Printf("The number of aInt is : %d\n", aInt)
	}
	fmt.Println("================ Cutting Line =================")
	bInt, err2 := strconv.Atoi("100")
	if err2 != nil {
		fmt.Printf("An error occured!%s\n", err2.Error())
		//状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
		os.Exit(1)
	} else {
		fmt.Printf("The number of aInt is : %d\n", bInt)
	}

	fmt.Println("================ Cutting Line =================")

	flag, err := openFile("if_demo2.go")
	if err != nil {
		fmt.Printf("An error occured!%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("The flag is : ", flag)

	fmt.Println("================ Cutting Line =================")

	if t, ok := mySqrt(25.0); ok {
		fmt.Println("t:", t)
	}
}

func openFile(fileName string) (flag bool, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}
	if err := file.Chmod(777); err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}
