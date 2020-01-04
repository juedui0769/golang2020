package ch02

import (
	"fmt"
	"runtime"
	"testing"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}

var info string // 非局部变量，string类型，未被显示赋值

// 和书上的例子不太一样，一开始还有点担心是否能够验证到`init()`的执行
// 结果担心是多余的，执行下面的测试func，init()是会执行的。
func Test01(t *testing.T){
	t.Log("-----")
	t.Log(info)
}