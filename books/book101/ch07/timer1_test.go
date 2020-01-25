package ch07

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer01(t *testing.T) {
	timer := time.NewTimer(time.Second * 2)
	now := time.Now()
	fmt.Printf("Now time: %v.\n", now)
	expire := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expire)
}



