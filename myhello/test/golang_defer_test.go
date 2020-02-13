package test

import (
	"fmt"
	"testing"
)

// 没想好怎么写这个测试，先看看书，再仔细体会下概念
func TestDefer01(t *testing.T) {
	deferA()
}

func deferA() {
	fmt.Println("hello defer")
}