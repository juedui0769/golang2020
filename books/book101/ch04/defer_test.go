package ch04

import (
	"errors"
	"fmt"
	"testing"
)

func isPositiveEvenNumber(number int) (result bool) {
	defer fmt.Println("done.")
	if number < 0{
		panic(errors.New("The number is a negative number!"))
	}
	if number %2==0{
		return true
	}
	return
}

// 运行下面的测试，三种情况下，"done" 都输出了，说明 `defer` 语句会在 `panic`,`return`之前执行！
func TestEvenNumber01(t *testing.T){
	t.Log(isPositiveEvenNumber(5))
	t.Log(isPositiveEvenNumber(8))
	t.Log(isPositiveEvenNumber(-1))
}