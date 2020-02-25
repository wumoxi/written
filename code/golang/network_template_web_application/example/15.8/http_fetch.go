package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://meihuahai.com")
	checkError(err)
	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(bytes))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
