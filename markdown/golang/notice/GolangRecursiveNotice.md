# Golang『递归』注意点

## 斐波那契数

### 使用递归获取指定数字的斐波那契数

```go
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%d) is: %d\n", i, fib(i))
	}
}

// fib获取斐波那契数
func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
```
程序输出0到10之间的斐波那契数如下

```shell
fib(0) is: 1
fib(1) is: 1
fib(2) is: 2
fib(3) is: 3
fib(4) is: 5
fib(5) is: 8
fib(6) is: 13
fib(7) is: 21
fib(8) is: 34
fib(9) is: 55
fib(10) is: 89
```

### 考虑如何重写上面的递归，并返回两个命名返回值，即数列中的位置和对应的值，例如 4 与 5，10 与 89。


```go
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		index, val := fib(i)
		fmt.Printf("fib(%d) is: %d\n", index, val)
	}
}

// fib获取斐波那契数
func fib(n int) (index, val int) {
	index = n
	if n <= 1 {
		val = 1
	} else {
		_, val1 := fib(n - 1)
		_, val2 := fib(n - 2)
		val = val1 + val2
	}
	return
}
```

程序输出如下

```shell
fib(0) is: 1
fib(1) is: 1
fib(2) is: 2
fib(3) is: 3
fib(4) is: 5
fib(5) is: 8
fib(6) is: 13
fib(7) is: 21
fib(8) is: 34
fib(9) is: 55
fib(10) is: 89
```

## 使用递归函数从 10 打印到 1

```go
package main

import "fmt"

func main() {
	out(10)
}

func out(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		n--
		out(n)
	}
}
```

程序输出如下

```shell
10 9 8 7 6 5 4 3 2 1 
```

## 阶乘

### 实现一个输出前 30 个整数的阶乘的程序。

n! 的阶乘定义为：_**`n! = n * (n-1)!`**_, _**`0! = 1`**_，因此它非常适合使用递归函数来实现。

```go
package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("factorial(%02d) is: %d\n", i, factorial(i))
	}
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}
```

程序运行结果如下，使用int64类型来计算阶乘，最多只能计算到20的阶乘，就溢出了。那解决这个问题，必须使用bigint来解决！

```shell
factorial(00) is: 1
factorial(01) is: 1
factorial(02) is: 2
factorial(03) is: 6
factorial(04) is: 24
factorial(05) is: 120
factorial(06) is: 720
factorial(07) is: 5040
factorial(08) is: 40320
factorial(09) is: 362880
factorial(10) is: 3628800
factorial(11) is: 39916800
factorial(12) is: 479001600
factorial(13) is: 6227020800
factorial(14) is: 87178291200
factorial(15) is: 1307674368000
factorial(16) is: 20922789888000
factorial(17) is: 355687428096000
factorial(18) is: 6402373705728000
factorial(19) is: 121645100408832000
factorial(20) is: 2432902008176640000
factorial(21) is: -4249290049419214848
factorial(22) is: -1250660718674968576
factorial(23) is: 8128291617894825984
factorial(24) is: -7835185981329244160
factorial(25) is: 7034535277573963776
factorial(26) is: -1569523520172457984
factorial(27) is: -5483646897237262336
factorial(28) is: -5968160532966932480
factorial(29) is: -7055958792655077376
```

使用bigint解决

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("factorial(%02d) is: %d\n", i, factorial(big.NewInt(int64(i))))
	}
}

func factorial(n *big.Int) *big.Int {
	if n.Int64() == 0 {
		return big.NewInt(1)
	}
	return n.Mul(n, factorial(big.NewInt(n.Int64()-1)))
}
```

程序运行结果如下

```go
factorial(00) is: 1
factorial(01) is: 1
factorial(02) is: 2
factorial(03) is: 6
factorial(04) is: 24
factorial(05) is: 120
factorial(06) is: 720
factorial(07) is: 5040
factorial(08) is: 40320
factorial(09) is: 362880
factorial(10) is: 3628800
factorial(11) is: 39916800
factorial(12) is: 479001600
factorial(13) is: 6227020800
factorial(14) is: 87178291200
factorial(15) is: 1307674368000
factorial(16) is: 20922789888000
factorial(17) is: 355687428096000
factorial(18) is: 6402373705728000
factorial(19) is: 121645100408832000
factorial(20) is: 2432902008176640000
factorial(21) is: 51090942171709440000
factorial(22) is: 1124000727777607680000
factorial(23) is: 25852016738884976640000
factorial(24) is: 620448401733239439360000
factorial(25) is: 15511210043330985984000000
factorial(26) is: 403291461126605635584000000
factorial(27) is: 10888869450418352160768000000
factorial(28) is: 304888344611713860501504000000
factorial(29) is: 8841761993739701954543616000000
```

## 目录
[Back](../GolangNotice.md)