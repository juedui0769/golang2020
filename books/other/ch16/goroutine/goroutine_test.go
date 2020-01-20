package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}



