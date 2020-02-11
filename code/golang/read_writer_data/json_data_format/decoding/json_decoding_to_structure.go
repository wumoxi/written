package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	b := []byte(`{"Name":"小花","Age":28,"Parents":["张三","兰花"]}`)

	var m FamilyMember
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Printf("json decoding error: %s\n", err)
		return
	}
	fmt.Printf("json decoded data structure: %v\n", m)

	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)

	encoder := json.NewEncoder(file)
	encoder.Encode()
}
