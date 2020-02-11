package main

import (
	"encoding/gob"
	"fmt"
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

	// Using an gob encoder:
	file, err := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open the file error:%s\n", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encoding gob data error: %s\n", err)
		return
	}
	fmt.Printf("encoding successfully!")
}
