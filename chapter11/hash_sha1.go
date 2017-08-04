package main

import (
	"crypto/sha1"
	"fmt"
)

func hashSHA1() {
	hash := sha1.New()
	//write string to writer
	b := []byte("Hello World")
	hash.Write(b)
	//append hash result,%x  hex print
	fmt.Printf("Result : %x\n", hash.Sum(b))
	//append hash result,%x  dec
	fmt.Printf("Result : %d\n", hash.Sum(b))

	hash.Reset()
	data := []byte("We shall overcome")
	n, err := hash.Write(data)
	if n != len(data) || err != nil {
		fmt.Printf("Hash write error: %d / %v", n, err)
	}

	checksum := hash.Sum(b)
	fmt.Printf("Result: %x\n", checksum)
}
