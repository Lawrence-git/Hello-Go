package main

import (
	"bytes"
	"fmt"
)

func byteBuffer() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello ")
	buffer.WriteString("World")
	fmt.Println(buffer.String())
}
