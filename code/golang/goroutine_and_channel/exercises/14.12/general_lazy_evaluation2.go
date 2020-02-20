package main

import "fmt"

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
	fibFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUint64Evaluator(fibFunc, []uint64{1, 1})
	for i := 0; i < 10; i++ {
		fmt.Printf("%vth fib: %v\n", i, fib())
	}
}

func BuildLazyUint64Evaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	go func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}()

	return func() Any {
		return <-retValChan
	}
}
