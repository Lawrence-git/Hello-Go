package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := [...]string{"http://www.jianshu.com/p/79caa1cc49a5",
		"http://www.jianshu.com/p/de44c0cc651a"}
	ch := make(chan string)
	for _, url := range urls {
		go fetchChannel(url, ch)
	}
	for range urls {
		fmt.Printf("%s\n", <-ch)
	}
}

func fetchChannel(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch err %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch err %v\n", err)
		os.Exit(1)
	}
	ch <- fmt.Sprintf("%s   %7d", url, nBytes)
}
