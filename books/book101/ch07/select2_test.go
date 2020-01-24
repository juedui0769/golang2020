package ch07

import (
	"fmt"
	"testing"
)

// 如果被执行的select语句中的所有case都不满足选择条件并且没有default case的话，
// 那么当前Goroutine就会一直被阻塞于此，直到某一个case中的发送或接收操作可以被立即进行为止。
// 注意，如果这样的select语句中的所有case右边的通道都是nil，那么当前Goroutine就会永远地被阻塞在这条select语句上！

// 所以，最好总是包含default case

// default case 可以被放置在该语句的任何位置上

func TestSelectDefaultCase01(t *testing.T) {
	ch10 := make(chan int, 10)
	ch10 <- 1
	select {
	default:
		fmt.Println("default")
	case e10 := <-ch10:
		fmt.Printf("1th case is selected. e10=%v.\n", e10)
	}
}
