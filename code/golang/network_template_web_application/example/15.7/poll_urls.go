package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.com",
	"http://alibaba.com/",
}

func main() {
	for _, url := range urls {
		response, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
			continue
		}
		fmt.Println(url, ": ", response.Status, response.Header)
	}
}
