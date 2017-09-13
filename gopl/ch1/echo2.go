package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo2Demo() {
	start := time.Now()
	var commandArgs []string
	for _, v := range os.Args[1:] {
		commandArgs = append(commandArgs, v)
	}
	fmt.Printf("%s\n", strings.Join(commandArgs, "\t"))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(elapsed.Nanoseconds()) //47251ns
}
