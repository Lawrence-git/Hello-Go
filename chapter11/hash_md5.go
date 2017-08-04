package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hasher := md5.New()
	byte := []byte("123456")
	hasher.Write(byte)
	fmt.Printf("Result: %x\n", hasher.Sum(byte))
	fmt.Printf("Result: %d\n", hasher.Sum(byte))
}
