package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rservice"
)

func main() {
	// 服务端注册rpc服务
	err := rpc.Register(rservice.TestService{})
	if err != nil {
		panic(err)
	}

	// 监听服务端口
	listener, err := net.Listen("tcp", ":8859")
	if err != nil {
		panic(err)
	}

	// 接收连接并处理服务调用
	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}
}
