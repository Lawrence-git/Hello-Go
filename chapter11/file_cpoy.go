package main

import (
	"fmt"
	"io"
	"os"
)

func fileCopy() {
	copyFile("demo1.txt", "target.txt")
	fmt.Println("Copy done!")
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err1 := os.Open(srcName)
	if err1 != nil {
		panic(err1)
	}
	defer src.Close()

	dst, err2 := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err2 != nil {
		panic(err2)
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
