package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op * 2
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func TestFn2(t *testing.T) {
	_func := timeSpent(slowFunc)
	t.Log(_func(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestFn3(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(2, 3, 5, 8))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("Start")
}
