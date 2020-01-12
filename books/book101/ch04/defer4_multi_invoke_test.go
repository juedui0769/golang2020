package ch04

import (
	"fmt"
	"testing"
)
/*
Possible resource leak, 'defer' is called in a for loop
Inspection info: Reports defer statements inside loops.
Using defer in loops can lead to resource leaks or unpredicatable execution order of the statements.
---
可能的资源泄漏，在for循环中调用“defer”
检查信息：报告在循环内延迟语句。
在循环中使用defer可能会导致资源泄漏或语句的执行顺序不可预测。
*/
/*
fmt.Printf("%d ", 0)
fmt.Printf("%d ", 1)
fmt.Printf("%d ", 2)
fmt.Printf("%d ", 3)
fmt.Printf("%d ", 4)
 */
func printNumbers() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func TestDefer401(t *testing.T) {
	printNumbers()
}




