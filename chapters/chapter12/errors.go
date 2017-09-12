package main

import (
	"fmt"
	"os"
)

func errorsDemo() {
	str, err := testError("Hello World")
	if err != nil {
		fmt.Printf("The error is %s!\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", str)

}

func testError(str string) (result string, err error) {
	if str == "Hello World" {
		// err = errors.New("error:Hello World is denied!")
		return result, fmt.Errorf("error:Hello World is denied!")
	}
	result = "done!"
	return result, nil
}
