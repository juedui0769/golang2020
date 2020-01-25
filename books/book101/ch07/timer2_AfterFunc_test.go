package ch07

import (
	"fmt"
	"testing"
	"time"
)

func TestTimerAfterFunc01(t *testing.T) {
	fmt.Println(">> do something ...")
	var timer *time.Timer
	f := func() {
		fmt.Printf("Expiration time: %v.\n", time.Now())
		fmt.Printf("C's len: %d\n", len(timer.C))
	}
	timer = time.AfterFunc(time.Second*1, f)
	time.Sleep(time.Second * 2)
}
