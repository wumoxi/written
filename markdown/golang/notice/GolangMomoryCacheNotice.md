# Golang通过『内存缓存』来提升性能

先来看一下通过递归来计算前50个斐波那契数列程序的运行性能！

```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 51

func main() {
	start := time.Now()
	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%d) is %d\n", i, fib(i))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %v\n", delta)
}

// fib计算斐波那契数
func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return res
}
```

程序的输出结果如下

```shell
fib(0) is 1
fib(1) is 1
fib(2) is 2
fib(3) is 3
fib(4) is 5
fib(5) is 8
fib(6) is 13
fib(7) is 21
fib(8) is 34
fib(9) is 55
fib(10) is 89
fib(11) is 144
fib(12) is 233
fib(13) is 377
fib(14) is 610
fib(15) is 987
fib(16) is 1597
fib(17) is 2584
fib(18) is 4181
fib(19) is 6765
fib(20) is 10946
fib(21) is 17711
fib(22) is 28657
fib(23) is 46368
fib(24) is 75025
fib(25) is 121393
fib(26) is 196418
fib(27) is 317811
fib(28) is 514229
fib(29) is 832040
fib(30) is 1346269
fib(31) is 2178309
fib(32) is 3524578
fib(33) is 5702887
fib(34) is 9227465
fib(35) is 14930352
fib(36) is 24157817
fib(37) is 39088169
fib(38) is 63245986
fib(39) is 102334155
fib(40) is 165580141
fib(41) is 267914296
fib(42) is 433494437
fib(43) is 701408733
fib(44) is 1134903170
fib(45) is 1836311903
fib(46) is 2971215073
fib(47) is 4807526976
fib(48) is 7778742049
fib(49) is 12586269025
fib(50) is 20365011074
time consuming: 4m20.716277099s
```

从结果可以看到这个计算过程需要将近4分钟20秒，好恐怖的样子！

_**注意：当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。最明显的例子就是生成斐波那契数列的程序。**_

_**要计算数列中第 n 个数字，需要先得到之前两个数的值，但很明显绝大多数情况下前两个数的值都是已经计算过的。即每个更后面的数都是基于之前计算结果的重复计算，正如上面这个示例展示的那样。**_

_**而我们要做就是将第 n 个数的值存在数组中索引为 n 的位置，然后在数组中查找是否已经计算过，如果没有找到，则再进行计算。**_

下面就通过利用『内存缓存』方式来提升程序的性能。


```go
package main

import (
	"fmt"
	"time"
)

// 计算前50位
const LIM = 51

// 斐波那契数数组
var fibs [LIM]int

func main() {
	start := time.Now()
	for i := 0; i < 51; i++ {
		fmt.Printf("fib(%d) is %d\n", i, fib(i))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time consuming: %d\n", delta)
}

// fib计算斐波那契数
func fib(n int) (res int) {
	// 已计算过
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}

	// 未计算过，计算过后放入到斐波那契数数组
	fibs[n] = res

	return res
}
```

程序输出如下

```shell
fib(0) is 1
fib(1) is 1
fib(2) is 2
fib(3) is 3
fib(4) is 5
fib(5) is 8
fib(6) is 13
fib(7) is 21
fib(8) is 34
fib(9) is 55
fib(10) is 89
fib(11) is 144
fib(12) is 233
fib(13) is 377
fib(14) is 610
fib(15) is 987
fib(16) is 1597
fib(17) is 2584
fib(18) is 4181
fib(19) is 6765
fib(20) is 10946
fib(21) is 17711
fib(22) is 28657
fib(23) is 46368
fib(24) is 75025
fib(25) is 121393
fib(26) is 196418
fib(27) is 317811
fib(28) is 514229
fib(29) is 832040
fib(30) is 1346269
fib(31) is 2178309
fib(32) is 3524578
fib(33) is 5702887
fib(34) is 9227465
fib(35) is 14930352
fib(36) is 24157817
fib(37) is 39088169
fib(38) is 63245986
fib(39) is 102334155
fib(40) is 165580141
fib(41) is 267914296
fib(42) is 433494437
fib(43) is 701408733
fib(44) is 1134903170
fib(45) is 1836311903
fib(46) is 2971215073
fib(47) is 4807526976
fib(48) is 7778742049
fib(49) is 12586269025
fib(50) is 20365011074
time consuming: 95.117µs
```

从结果可以看出，这个使用了『内存缓存』方式的优化程序，它仅仅只需要 `95.117µs` 微秒，根据公式 `1minute=60 000 000µs`，之前的程序运行时间为4分20秒，那么就提升了250万倍，结果惊人，而且系统资源的使用也可以忽略不计。

内存缓存的技术在使用计算成本相对昂贵的函数时非常有用（不仅限于例子中的递归），譬如大量进行相同参数的运算。这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数。

## 目录
[Back](../GolangNotice.md)