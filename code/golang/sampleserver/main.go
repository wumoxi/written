package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"username": "zhangsan", "age": "18", "sex": "male", "mobile": "13927928289"}`))
	})
	log.Fatal(http.ListenAndServe(":8859", nil))
}
