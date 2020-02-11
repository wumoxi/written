package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	// Initialized json decoding storage variable
	var v VCord

	// Read all file contents
	data, err := ioutil.ReadFile("code/golang/read_writer_data/json_data_format/vcard.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read all file contents error: %s\n", err)
		return
	}

	// JSON decoding(JSON解码)
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "json decoding error: %s\n", err)
		return
	}

	fmt.Printf("json decoded data structure: %v\n", v)
}
