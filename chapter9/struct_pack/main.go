package main

import (
	"fmt"
	"reflect"

	custom "./custom_pack"
)

func main() {
	person := custom.Person{"Xia", 11, "HangZhou"}
	fmt.Printf("%v\n", person)

	fmt.Println("===========================")

	for i := 0; i < 3; i++ {
		refPerson(person, i)
	}
}

func refPerson(person custom.Person, index int) {
	refType := reflect.TypeOf(person)
	indexField := refType.Field(index)
	fmt.Println(indexField.Tag)
}
