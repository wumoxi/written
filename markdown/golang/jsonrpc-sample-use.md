# jsonRPC简单使用

目录结构如下：

```shell
rservice
├── rservice.go
├── client
│   └── client.go
└── server
    └── server.go

```

## 定义RPC服务方法

```go
// rservice.go
package rservice

import "errors"

// 定义rpc服务
type TestService struct {}

// 定义rpc服务参数结构类型
type Args struct {
	A, B int
}

// 定义rpc服务方法
func (TestService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
```

上面定义了一个名为 `TestService` 的服务，这个服务有一个方法 `Div`，这个方法就是进行简单的业务处理也就是进行除法运算，并把运算结果写入到 `result` 变量内。

## 注册RPC服务

编写RPC服务端服务程序

```go
// server/server.go
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
```

启动RPC服务端服务程序

```shell
$ go run server/server.go
```

## 调用RCP服务

### 使用telnet工具进行RPC服务调用(手动调用)

```shell
$ telnet 127.0.0.1 8859
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```

输入调用参数

```shell
{"method":"TestService.Div", "params":[{"A":15, "B":6}], "id": 12345}
```

- method参数为要调用的RPC服务方法
- params参数为调用RPC服务方法需要传递的参数
- id参数为调用编号，调用结束后服务端会原样返回

----------------------------------------

返回调用结果

```shell
{"id":12345,"result":2.5,"error":null}
```

### 使用RPC客户端进行RPC服务调用(自动调用)

编写调用程序

```go
// client/client.go
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

	// 声明RPC调用结果
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
```

返回调用结果

```shell
RPC call result: 2.222222
RPC call error: division by zero
```
