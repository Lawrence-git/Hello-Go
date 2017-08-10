package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host   = "zhihu.com"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		flag   = true
		count  = 0
	)
	// 创建一个socket
	con, err := net.Dial("tcp", remote)
	// 发送我们的消息，一个http GET请求
	io.WriteString(con, msg)
	// 读取服务器的响应
	for flag {
		count, err = con.Read(data)
		flag = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
}
