package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "My name is %s,age is %d!\n", "Xia", 23)

	//缓冲区输出
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(writer, "My name is %s,age is %d!\n", "Xia", 23)
	writer.Flush()
}
