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
