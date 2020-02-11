package main

import (
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

func main() {
	// Open the file. (打开文件)
	file, err := os.Open("vcard.gob")
	if err != nil {
		log.Fatalf("open the file error: %s\n", err)
	}
	// Close the open file. (关闭打开的文件)
	defer file.Close()

	// Decode the gob content to structure. (将gob内容解码为结构)
	var vc VCard
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&vc)
	if err != nil {
		log.Fatalf("exercises the error: %s\n", err)
	}

	// Print the decoded structure data. (打印解码后的结构数据)
	fmt.Printf("exercises data structure: %+v\n", vc)
}
