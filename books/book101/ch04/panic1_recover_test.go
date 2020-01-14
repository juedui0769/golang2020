package ch04

import (
	"fmt"
	"testing"
)

func fetchElement(ss []string, index int) (element string) {
	if index >= len(ss) {
		fmt.Printf("Occur a panic! [index=%d]\n", index)
		panic(index)
	}
	fmt.Printf("Fetching the elment... [index=%d]\n", index)
	element = ss[index]
	defer fmt.Printf("The elment is \"%s\". [index=%d]\n", element, index)
	fetchElement(ss, index+1)
	return
}

func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered a panic. [index=%d]\n", v)
		}
	}()
	ss := []string{"A", "B", "C"}
	fmt.Printf("Fetch the elments in %v one by one...\n", ss)
	fetchElement(ss, 0)
	fmt.Println("The elements fetching is done.")
}

func TestPanicRecover01(t *testing.T) {
	fetchDemo()
	fmt.Println("The test function is executed.")
}