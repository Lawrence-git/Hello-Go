package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	listener, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}
