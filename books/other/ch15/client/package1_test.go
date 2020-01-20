package client

import (
	"books/other/ch15/package"
	"testing"
)

func TestPackage001(t *testing.T) {
	t.Log(_package.GetFib(5))
	t.Log(_package.Square(5))
}
