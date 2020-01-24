package ch07

import (
	"fmt"
	"testing"
)

// 当发现第一个满足选择条件的case时，运行时系统就会执行该case所包含的语句。
// 这也意味着其他case会被忽略。
// 如果同时有多个case满足条件，那么运行时系统会通过一个伪随机的算法决定哪一个case将会被执行。
// 例如，下面的代码会随机的向一个通道发送5个范围在[1,3]的整数：

func TestSelect01(t *testing.T) {
	chanCap := 5
	ch7 := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 2:
		case ch7 <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch7)
	}
}




