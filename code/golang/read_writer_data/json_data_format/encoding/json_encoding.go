package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCord struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	// Structure source data(源数据结构)
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCord{FirstName: "Jan", LastName: "Kersschot", Addresses: []*Address{pa, wa}, Remark: "none"}
	// fmt.Printf("%v\n", vc) // {Jan Kersschot [0xc000062180 0xc0000621b0] none}

	// JSON encoding(JSON编码):
	js, err := json.Marshal(vc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %s\n", err)
	}
	fmt.Printf("JSON encoding after content: %s\n", js)

	// Using an encoder(使用编码器将结构数据编码后写入到文件)
	file, err := os.OpenFile("code/golang/read_writer_data/json_data_format/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open the file error: %s\n", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encoding to file error: %s\n", err)
	}
	fmt.Println("Processing complete!")
}
