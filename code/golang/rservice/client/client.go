package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"rservice"
)

func main() {
	// 连接到RPC服务端
	conn, err := net.Dial("tcp", ":8859")
	if err != nil {
		panic(err)
	}

	// 创建RPC客户端
	client := jsonrpc.NewClient(conn)

	// 声明RPC调用回复结果
	var reply float64

	// 调用RPC服务
	err = client.Call("TestService.Div", rservice.Args{A:20, B:9}, &reply)
	if err != nil {
		fmt.Printf("RPC call error: %s\n", err)
	} else {
		fmt.Printf("RPC call result: %f\n", reply)
	}

	err = client.Call("TestService.Div", rservice.Args{A:20, B:0}, &reply)
	if err != nil {
		fmt.Printf("RPC call error: %s\n", err)
	} else {
		fmt.Printf("RPC call result: %f\n", reply)
	}
}
