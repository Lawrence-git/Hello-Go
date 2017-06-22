package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
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
