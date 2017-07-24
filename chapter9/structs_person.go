package main

import (
	"fmt"
	"strings"
)

//Person structs demo
type Person struct {
	firstName string
	lastName  string
	age       int
}

//传引用，可以修改实际对象的值
func updatePersonByReference(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
	p.age = p.age * 2
}

//传值，不修改实际对象的值
func updatePersonByCopy(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
	p.age = p.age * 2
}

func strucsPerson() {
	var pers1 *Person
	pers1 = new(Person)
	pers1.firstName = "Xia"
	pers1.lastName = "Zhengyue"
	pers1.age = 23
	fmt.Printf("The name of the person is %s %s.The age is %d.\n",
		pers1.firstName, pers1.lastName, pers1.age)
	updatePersonByReference(pers1)
	fmt.Printf("The name of the person is %s %s.The age is %d.\n",
		pers1.firstName, pers1.lastName, pers1.age)

	fmt.Println()

	var pers2 = Person{"Ce", "Shi", 111}
	fmt.Printf("The name of the person is %s %s.The age is %d.\n",
		pers2.firstName, pers2.lastName, pers2.age)
	updatePersonByCopy(pers2)
	fmt.Printf("The name of the person is %s %s.The age is %d.\n",
		pers2.firstName, pers2.lastName, pers2.age)
	updatePersonByReference(&pers2)
	fmt.Printf("The name of the person is %s %s.The age is %d.\n",
		pers2.firstName, pers2.lastName, pers2.age)
}
