package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	if i, ok := p.(int); ok{
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok{
		fmt.Println("String", s)
		return
	}
	fmt.Println("Unkonw Type")
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello world")
}
