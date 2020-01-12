package ch04

import "testing"

func appendNumbers(ints []int) (result []int) {
	result = append(ints, 1)
	defer func() {
		result = append(result, 2)
	}()
	result = append(result, 3)
	defer func() {
		result = append(result, 4)
	}()
	result = append(result, 5)
	defer func() {
		result = append(result, 6)
	}()
	return result
}

/**
[0 1 3 5 6 4 2]
根据结果观察调用顺序
*/
func TestDefer5RunOrder(t *testing.T) {
	result := appendNumbers([]int{0})
	t.Log(result)
}

