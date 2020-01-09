package ch04

import "testing"

// 像 if

func TestSwitchLikeIf001(t *testing.T) {
	switch number := 123; {
	case number < 100:
		number++
	case number < 200:
		number--
	default:
		number -= 2
	}
}

