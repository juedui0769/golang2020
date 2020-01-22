package ch07

import (
	"fmt"
	"runtime"
	"testing"
)

/*
	(1) 如果不传入name，而直接使用，输出接结果将全部是对"Mark"的问候。
	(2) `runtime.Gosched` 会让出运行权，让其他Goroutine先执行
 */
func TestHelloGoroutine01(t *testing.T) {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s.\n", who)
		}(name)
	}
	runtime.Gosched()
}






