package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8088")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("First, what is your name?")
		clientName, _ := inputReader.ReadString('\n')
		// fmt.Printf("CLIENTNAME %s", clientName)
		// Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
		// trimmedClient := strings.Trim(clientName, "\n")
		_, err = conn.Write([]byte(clientName))
		if err != nil {
			fmt.Println("Error Writing", err.Error())
			return // 终止程序
		}
	}
}
