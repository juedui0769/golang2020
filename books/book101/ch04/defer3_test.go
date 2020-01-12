package ch04

import (
	"fmt"
	"testing"
)

func begin(funcName string) string {
	fmt.Printf(">> Enter function %s.\n", funcName)
	return funcName
}

func end(funcName string) string {
	fmt.Printf("<< Exit function %s.\n", funcName)
	return funcName
}

func record() {
	defer end(begin("record"))
	fmt.Println("In function record.")
}

func TestDefer03(t *testing.T) {
	record()
}

