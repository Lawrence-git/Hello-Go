package main

import (
	"fmt"
	"os"
	"time"
)

func echo1Demo() {
	start := time.Now()
	var commandArgs string
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		commandArgs += v + "|"
	}
	fmt.Printf("%s\n", commandArgs)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed.Nanoseconds()) //52713ns
}
