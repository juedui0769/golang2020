package ch04

import "testing"

func TestFor01(t *testing.T) {
	number := 1
	for number < 200 {
		number += 2
	}
}

func TestFor02(t *testing.T) {
	number := 1
	for {
		number++
		// 如果没有下面的break语句，这就是一个死循环。
		if number == 200 {
			break
		}
	}
}
