package hello

import (
	"errors"
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("hello,world!")
	s := "hello"
	c := []byte(s)
	t.Log(c)
	c[0] = 'c'
	s2 := string(c)
	t.Log(s2)
}

func TestError(t *testing.T)  {
	err := errors.New("错误")
	if err != nil {
		t.Log(err)
	}
}

