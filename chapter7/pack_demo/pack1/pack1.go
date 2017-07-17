// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Package pack1 demo
package pack1

import "fmt"

//Pack1Int  Int Demo
var Pack1Int = 42

//PackFloat Float Demo
var PackFloat = 3.14

func init() {
	fmt.Println("pack1.go init")
}

//ReturnStr return a string
func ReturnStr() string {
	return "Hello main!Here is package pack1"
}
