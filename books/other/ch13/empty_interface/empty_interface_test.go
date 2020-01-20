package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("String", s)
	//	return
	//}
	//fmt.Println("Unkown Type")


	// fmt.Println(p.(type))
	// 上面的语句不能直接打印, 否则会报错
	// Use of .(type) outside type switch

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unkown Type")
	}

}

func TestEmptyInterface01(t *testing.T) {
	DoSomething(1)
	DoSomething("golang")
	DoSomething(1.23)
}
