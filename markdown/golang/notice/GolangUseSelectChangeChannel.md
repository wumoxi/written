# Golang『使用Select切换协程』注意点

## 使用Select计算fibonacci数

请看下面一个使用Select和多通道计算fibonacci数据的经典案例：

```go
package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		// Loop read data for the channel
		for i := 0; i < 30; i++ {
			fmt.Printf("fibonacci(%d) is: %d\n", i, <-ch)
		}

		// Write data to the channel
		done <- true
	}()
	fibonacci(ch, done)
}

func fibonacci(c chan int, done chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // send data to the channel
			x, y = y, x+y
		case <-done: // read data for the channel
			fmt.Println("Send data to the channel is done!")
			return
		}
	}
}
```

运行程序可以得到如下所示结果：

```text
fibonacci(0) is: 1
fibonacci(1) is: 1
fibonacci(2) is: 2
fibonacci(3) is: 3
fibonacci(4) is: 5
fibonacci(5) is: 8
fibonacci(6) is: 13
fibonacci(7) is: 21
fibonacci(8) is: 34
fibonacci(9) is: 55
fibonacci(10) is: 89
fibonacci(11) is: 144
fibonacci(12) is: 233
fibonacci(13) is: 377
fibonacci(14) is: 610
fibonacci(15) is: 987
fibonacci(16) is: 1597
fibonacci(17) is: 2584
fibonacci(18) is: 4181
fibonacci(19) is: 6765
fibonacci(20) is: 10946
fibonacci(21) is: 17711
fibonacci(22) is: 28657
fibonacci(23) is: 46368
fibonacci(24) is: 75025
fibonacci(25) is: 121393
fibonacci(26) is: 196418
fibonacci(27) is: 317811
fibonacci(28) is: 514229
fibonacci(29) is: 832040
Send data to the channel is done!
```

在这个示例中我们可以看到在select通道多路复用中，既有数据发送操作也有数据接收操作哦，可以仔细揣摩这种用法哦，多多有益哦！

## 使用Select生成随机位

做一个随机位生成器，程序可以提供无限的随机 0 或者 1 的序列

```go
package main

import "fmt"

func main() {
	c := make(chan int)

	// consumer
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()

	// producer
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
```

运行程序输出如下结果：

```shell
1 1 1 0 1 1 0 1 1 0 0 0 0 1 1 0 1 0 1 1 1 1 0 1 0 1 1 1 1 0 0 1 
```

## 目录
[Back](../GolangNotice.md)