package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}
	fmt.Println("Unkown Type")
}

func TestEmptyInterface01(t *testing.T) {
	DoSomething(1)
	DoSomething("golang")
	DoSomething(1.23)
}
