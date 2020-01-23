package ch07

import (
	"fmt"
	"testing"
	"time"
)

// 在发送端调用`close`以关闭通道却不会对接收端接收该通道中已有的元素值产生任何影响。

func TestChanClose01(t *testing.T) {
	chInt := make(chan int, 5)
	sign := make(chan byte, 2)

	go func() {
		for i := 0; i < 5; i++ {
			chInt <- i
			time.Sleep(time.Second * 1)
		}
		close(chInt)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()
	go func() {
		for {
			e, ok := <-chInt
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		fmt.Println(">> Done.")
		sign <- 1
	}()
	<-sign
	<-sign
}


