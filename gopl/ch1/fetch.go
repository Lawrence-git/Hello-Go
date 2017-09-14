package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	fetchURL = "http://www.jianshu.com/p/79caa1cc49a5"
)

func main() {
	resp, _ := http.Get(fetchURL)
	content, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", content)
}
