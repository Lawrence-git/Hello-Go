package main

import (
	"fmt"
	"reflect"

	custom "./custom_pack"
)

func main() {
	person := new(custom.Person)
	person.SetName("Xia")
	person.SetAge(23)
	person.SetAddress("HangZhou")
	fmt.Printf("%v\n", person)
	fmt.Printf("The name is %s. The age is %d. The address is %s",
		person.GetName(), person.GetAge(), person.GetAddress())

}

func refPerson(person custom.Person, index int) {
	refType := reflect.TypeOf(person)
	indexField := refType.Field(index)
	fmt.Println(indexField.Tag)
}
