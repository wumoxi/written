# Golang『big package』注意点

## 大整数运算

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 这是bigInts的一些计算：
	ia := big.NewInt(1996)
	ib := big.NewInt(2)

	// 整型乘法运算
	product := ia.Mul(ia, ib)
	fmt.Printf("big operation: ia * ib = %d, ia = %d, ib = %d\n", product, ia, ib)

	// 整型加法运算
	sum := ia.Add(ia, ib)
	fmt.Printf("big operation: ia + ib = %d, ia = %d, ib = %d\n", sum, ia, ib)

	// 整型减法运算
	sub := ia.Sub(ia, ib.Add(ib, big.NewInt(98)))
	fmt.Printf("big operation: ia - ib = %d, ia = %d, ib = %d\n", sub, ia, ib)

	// 整型除法运算
	div := ia.Div(ia, ib)
	fmt.Printf("big operation: ia / ib = %d, ia = %d, ib = %d\n", div, ia, ib)
}
```

程序输出如下

```shell
big operation: ia * ib = 3992, ia = 3992, ib = 2
big operation: ia + ib = 3994, ia = 3994, ib = 2
big operation: ia - ib = 3894, ia = 3894, ib = 100
big operation: ia / ib = 38, ia = 38, ib = 100
```

## 目录
[Back](../GolangNotice.md)