package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func getData(ch chan string) {
	for {
		fmt.Printf("%s ", <-ch)
	}
}

func sendData(ch chan string) {
	// 数据流向通道
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}
