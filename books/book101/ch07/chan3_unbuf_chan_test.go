package ch07

import (
	"fmt"
	"testing"
	"time"
)

func TestUnbufChan01(t *testing.T) {
	unbufChan := make(chan int)
	go func() {
		fmt.Println(">> Sleep a second...")
		time.Sleep(time.Second)
		num := <-unbufChan
		fmt.Printf(">>> Received a integer %d.\n", num)
	}()
	num := 1
	fmt.Printf("Send integer %d...\n", num)
	unbufChan <- num
	fmt.Println("Done.")
}

