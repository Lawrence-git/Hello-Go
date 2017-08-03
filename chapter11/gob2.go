// gob2.go
package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	//将vcard以gob序列化到文件中，然后在反序列化解析出来
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
		return
	}

	//反序列化
	var vcard VCard
	vcFile, _ := os.Open("vcard.gob")
	defer vcFile.Close()

	reader := bufio.NewReader(vcFile)
	dec := gob.NewDecoder(reader)
	error := dec.Decode(&vcard)
	if error != nil {
		log.Println("Error in decoding gob", error)
		return
	}
	fmt.Printf("%v\n,%v\t%v\n", &vcard, *(&vcard.Addresses[0]), *(&vcard.Addresses[1]))
}
