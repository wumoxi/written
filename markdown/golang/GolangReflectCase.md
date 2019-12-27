# Golang关于反射的使用案例

## 获取一个运行时的函数名称

可以使用`runtime.FuncForPC(uintptr).Name()`方法结合`reflect.ValueOf(fn).Pointer()`方法来实现！`reflect.ValueOf(fn).Pointer()`会以以`uintptr`的形式返回v的值。

```go
// GetFuncName 获取函数名称
func GetFuncName(fn interface{}) (string, error) {
	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), nil
	}
	return "", errors.New("fn parameter is not func type")
}
```

可以这么具体来使用

```go
package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	funcName, err := GetFuncName(main)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("GetFuncName(main): %s\n", funcName)
	}

	funcName, err = GetFuncName(fmt.Printf)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("GetFuncName(fmt.Printf): %s\n", funcName)
	}
}

// GetFuncName 获取函数名称
func GetFuncName(fn interface{}) (string, error) {
	if reflect.ValueOf(fn).Kind() == reflect.Func {
		return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), nil
	}
	return "", errors.New("fn parameter is not func type")
}
```

程序结果如下

```shell
GetFuncName(main): main.main
GetFuncName(fmt.Printf): fmt.Printf
```


## 目录
[Back](../../README.md)