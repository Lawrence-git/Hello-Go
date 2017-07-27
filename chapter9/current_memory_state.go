package main

import (
	"fmt"
	"runtime"
)

func main() {
	var status runtime.MemStats
	runtime.ReadMemStats(&status)
	fmt.Printf("%d Kb\n", status.Alloc/1024)
}
