package error

import (
	"errors"
	"testing"
)

var LessThanTwoError error = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}
		if err == LargerThanHundredError {
			t.Error("Need a small number")
		}
	} else {
		t.Log(v)
	}
	if v, err := GetFibonacci(10); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}
		if err == LargerThanHundredError {
			t.Error("Need a small number")
		}
	} else {
		t.Log(v)
	}
	if v, err := GetFibonacci(101); err != nil {
		if err == LessThanTwoError {
			t.Error("Need a larger number")
		}
		if err == LargerThanHundredError {
			t.Error("Need a small number")
		}
	} else {
		t.Log(v)
	}
}
