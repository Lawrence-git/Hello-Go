//针对标准输入的内容的去重复
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "q" {
			break
		}
		counts[input.Text()]++
	}
	for line, v := range counts {
		if v > 0 {
			fmt.Println(line)
		}
	}
}
