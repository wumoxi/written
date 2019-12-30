# Golang『斐波那契』多种实现方式

使用多种方式实现打印前50个Fibonacci数字程序

## 使用递归实现

```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 50

func main() {
	start := time.Now()
	fmt.Println("using the recursion instance fibonacci!")
	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%02d) is: %d\n", i, fib(i))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %v\n", delta)
}

// fib计算斐波那契数
func fib(LIM int) int {
	if LIM <= 1 {
		return 1
	}
	return fib(LIM-1) + fib(LIM-2)
}
```

程序输出如下

```shell
using the recursion instance fibonacci!
fib(00) is: 1
fib(01) is: 1
fib(02) is: 2
fib(03) is: 3
fib(04) is: 5
fib(05) is: 8
fib(06) is: 13
fib(07) is: 21
fib(08) is: 34
fib(09) is: 55
fib(10) is: 89
fib(11) is: 144
fib(12) is: 233
fib(13) is: 377
fib(14) is: 610
fib(15) is: 987
fib(16) is: 1597
fib(17) is: 2584
fib(18) is: 4181
fib(19) is: 6765
fib(20) is: 10946
fib(21) is: 17711
fib(22) is: 28657
fib(23) is: 46368
fib(24) is: 75025
fib(25) is: 121393
fib(26) is: 196418
fib(27) is: 317811
fib(28) is: 514229
fib(29) is: 832040
fib(30) is: 1346269
fib(31) is: 2178309
fib(32) is: 3524578
fib(33) is: 5702887
fib(34) is: 9227465
fib(35) is: 14930352
fib(36) is: 24157817
fib(37) is: 39088169
fib(38) is: 63245986
fib(39) is: 102334155
fib(40) is: 165580141
fib(41) is: 267914296
fib(42) is: 433494437
fib(43) is: 701408733
fib(44) is: 1134903170
fib(45) is: 1836311903
fib(46) is: 2971215073
fib(47) is: 4807526976
fib(48) is: 7778742049
fib(49) is: 12586269025
time consuming: 2m23.656896714s
```

## 闭包实现

```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 50

func main() {
	start := time.Now()
	fmt.Println("using the closure instance fibonacci!")
	fb := fib()
	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%02d) is: %d\n", i, fb())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %v\n", delta)
}

// fib计算斐波那契数
func fib() func() int {
	a, b := 1, 1
	return func() int {
		c := a
		a, b = b, b+c
		return c
	}
}
```

程序输出如下

```shell
using the closure instance fibonacci!
fib(00) is: 1
fib(01) is: 1
fib(02) is: 2
fib(03) is: 3
fib(04) is: 5
fib(05) is: 8
fib(06) is: 13
fib(07) is: 21
fib(08) is: 34
fib(09) is: 55
fib(10) is: 89
fib(11) is: 144
fib(12) is: 233
fib(13) is: 377
fib(14) is: 610
fib(15) is: 987
fib(16) is: 1597
fib(17) is: 2584
fib(18) is: 4181
fib(19) is: 6765
fib(20) is: 10946
fib(21) is: 17711
fib(22) is: 28657
fib(23) is: 46368
fib(24) is: 75025
fib(25) is: 121393
fib(26) is: 196418
fib(27) is: 317811
fib(28) is: 514229
fib(29) is: 832040
fib(30) is: 1346269
fib(31) is: 2178309
fib(32) is: 3524578
fib(33) is: 5702887
fib(34) is: 9227465
fib(35) is: 14930352
fib(36) is: 24157817
fib(37) is: 39088169
fib(38) is: 63245986
fib(39) is: 102334155
fib(40) is: 165580141
fib(41) is: 267914296
fib(42) is: 433494437
fib(43) is: 701408733
fib(44) is: 1134903170
fib(45) is: 1836311903
fib(46) is: 2971215073
fib(47) is: 4807526976
fib(48) is: 7778742049
fib(49) is: 12586269025
time consuming: 81.245µs
```

## 数组实现

```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 50

func main() {
	start := time.Now()
	fmt.Println("using the array instance fibonacci!")
	fibs := fib()
	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%02d) is: %d\n", i, fibs[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %v\n", delta)
}

// fib计算斐波那契数
func fib() [LIM]int {
	var arr = [LIM]int{}
	for i := 0; i < LIM; i++ {
		if i == 0 || i == 1 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}
```

程序输出如下

```shell
using the array instance fibonacci!
fib(00) is: 1
fib(01) is: 1
fib(02) is: 2
fib(03) is: 3
fib(04) is: 5
fib(05) is: 8
fib(06) is: 13
fib(07) is: 21
fib(08) is: 34
fib(09) is: 55
fib(10) is: 89
fib(11) is: 144
fib(12) is: 233
fib(13) is: 377
fib(14) is: 610
fib(15) is: 987
fib(16) is: 1597
fib(17) is: 2584
fib(18) is: 4181
fib(19) is: 6765
fib(20) is: 10946
fib(21) is: 17711
fib(22) is: 28657
fib(23) is: 46368
fib(24) is: 75025
fib(25) is: 121393
fib(26) is: 196418
fib(27) is: 317811
fib(28) is: 514229
fib(29) is: 832040
fib(30) is: 1346269
fib(31) is: 2178309
fib(32) is: 3524578
fib(33) is: 5702887
fib(34) is: 9227465
fib(35) is: 14930352
fib(36) is: 24157817
fib(37) is: 39088169
fib(38) is: 63245986
fib(39) is: 102334155
fib(40) is: 165580141
fib(41) is: 267914296
fib(42) is: 433494437
fib(43) is: 701408733
fib(44) is: 1134903170
fib(45) is: 1836311903
fib(46) is: 2971215073
fib(47) is: 4807526976
fib(48) is: 7778742049
fib(49) is: 12586269025
time consuming: 80.032µs
```

## 切片实现

```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 50

func main() {
	start := time.Now()
	fmt.Println("using the slice instance fibonacci!")
	arr := generateSectionIntSliceOfOrderly(0, LIM, 1)
	fibs := fib(arr)
	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%02d) is: %d\n", i, fibs[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %v\n", delta)
}

// generateSectionIntSliceOfOrderly生成指定范围内的有序整型切片
func generateSectionIntSliceOfOrderly(min, max int, step int) []int {
	result := make([]int, max, max)
	for i := min; i < max; i += step {
		result[i] = i
	}
	return result
}

// fib计算斐波那契数
func fib(a []int) (fibs []int) {
	fibs = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 || i == 1 {
			fibs[i] = 1
		} else {
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
	}
	return
}
```

程序输出如下

```shell
using the slice instance fibonacci!
fib(00) is: 1
fib(01) is: 1
fib(02) is: 2
fib(03) is: 3
fib(04) is: 5
fib(05) is: 8
fib(06) is: 13
fib(07) is: 21
fib(08) is: 34
fib(09) is: 55
fib(10) is: 89
fib(11) is: 144
fib(12) is: 233
fib(13) is: 377
fib(14) is: 610
fib(15) is: 987
fib(16) is: 1597
fib(17) is: 2584
fib(18) is: 4181
fib(19) is: 6765
fib(20) is: 10946
fib(21) is: 17711
fib(22) is: 28657
fib(23) is: 46368
fib(24) is: 75025
fib(25) is: 121393
fib(26) is: 196418
fib(27) is: 317811
fib(28) is: 514229
fib(29) is: 832040
fib(30) is: 1346269
fib(31) is: 2178309
fib(32) is: 3524578
fib(33) is: 5702887
fib(34) is: 9227465
fib(35) is: 14930352
fib(36) is: 24157817
fib(37) is: 39088169
fib(38) is: 63245986
fib(39) is: 102334155
fib(40) is: 165580141
fib(41) is: 267914296
fib(42) is: 433494437
fib(43) is: 701408733
fib(44) is: 1134903170
fib(45) is: 1836311903
fib(46) is: 2971215073
fib(47) is: 4807526976
fib(48) is: 7778742049
fib(49) is: 12586269025
time consuming: 85.711µs
```