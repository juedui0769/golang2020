package ch04

import (
	"fmt"
	"testing"
)

func TestAnonymousDeferFunc01(t *testing.T) {
	defer func() {
		fmt.Println("The finishing touches.")
	}()
	fmt.Println("do something...")
}