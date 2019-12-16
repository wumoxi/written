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
