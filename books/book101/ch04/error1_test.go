package ch04

import (
	"fmt"
	"os"
	"testing"
)

func TestError01(t *testing.T) {
	file, err := os.Open("/etc/profile")
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			fmt.Printf("Path Error: %s (op=%s, path=%s)\n", pe.Err, pe.Op, pe.Path)
		} else {
			fmt.Printf("Unknown Error: %s\n", err)
		}
	} else {
		// do something with file
		_ = file
	}
}
