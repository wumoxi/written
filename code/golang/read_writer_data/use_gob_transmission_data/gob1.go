package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X    int
	Y    int
	Z    int
	Name string
}

type Q struct {
	X    *int32
	Y    *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.
	// 初始化编码器和解码器。
	// Normally enc and dec would be bound to network connections and the encoder and decoder would run in different processes.
	// 通常enc和dec是绑定到网络连接，编码器和解码器将在不同的进程中运行。

	// Stand-in for a network connection. (代替网络连接)
	var network bytes.Buffer

	// Will write to network. (将写入网络)
	enc := gob.NewEncoder(&network)

	// Will read from network. (将从网络读取)
	dec := gob.NewDecoder(&network)

	// Encode (send) the value. (编码(发送)值)
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatalf("encode error: %s\n", err)
	}

	// Decode (receive) the value. (解码(接收)值)
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatalf("decode error: %s\n", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
