package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntFunc func(op int) int

func timeSpent(inner IntFunc) IntFunc{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op * 2
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}