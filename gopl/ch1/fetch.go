package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	fetchURL = "www.jianshu.com/p/79caa1cc49a5"
)

func fetchMain() {
	url := fetchURL
	if !strings.HasPrefix(fetchURL, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch err %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	//ioutil.ReadAll 创建一个缓冲区来存储
	// content, err := ioutil.ReadAll(resp.Body)

	//将resp的body内容copy到标准输出中
	io.Copy(os.Stdout, resp.Body)
	//输出resp的状态码
	fmt.Printf("\n\n%d\t%s\n", resp.StatusCode, resp.Status)
}
