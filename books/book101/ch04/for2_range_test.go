package ch04

import (
	"fmt"
	"testing"
)

func TestForRange01(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	for i, d := range numbers {
		fmt.Printf("(Index): %d, (Value): %d\n", i, d)

	}
}
