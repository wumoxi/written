package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func getData(ch chan string) {
	for value := range ch {
		fmt.Println(value)
	}
}

func sendData(ch chan string) {
	ch <- "张三"
	ch <- "李四"
	ch <- "王五"
	ch <- "赵六"
	close(ch)
}
