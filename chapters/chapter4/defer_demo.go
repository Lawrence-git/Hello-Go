package main

import (
	"fmt"
	"time"
)

func deferDemo() {
	start := time.Now()
	getConnection()
	defer closeConnection()
	operateDBData()
	end := time.Now()
	// var temp = 0
	// for i := 0; i < 100000; i++ {
	// 	for j := 0; j < 100000; j++ {
	// 		temp = temp + 1
	// 	}
	// }
	fmt.Println(end.Sub(start))
}

func getConnection() {
	fmt.Println("Get connection!")
}

func closeConnection() {
	fmt.Println("Close connection!")
}

func operateDBData() {
	fmt.Println("Operate Database data!")
}
