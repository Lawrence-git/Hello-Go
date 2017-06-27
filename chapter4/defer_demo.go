package main

import "fmt"

func deferDemo() {
	getConnection()
	defer closeConnection()
	operateDBData()
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
