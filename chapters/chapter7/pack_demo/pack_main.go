package main

import (
	"fmt"

	pack1 "./pack1"
	_ "./pack2" //只执行pack2下的init方法
)

func main() {
	fmt.Println(pack1.Pack1Int)
	fmt.Println(pack1.PackFloat)
	fmt.Println(pack1.ReturnStr())
	fmt.Println(pack1.ReturnAnotherStr())
}
