package ch04

import (
	"fmt"
	"testing"
)

func TestForRange01(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	for i, d := range numbers {
		fmt.Printf("(Index): %d, (Value): %d\n", i, d)

	}
}

// `namesCount` 包含了某个网站的所有用户昵称及其重复次数
// 现在我们想从中查找到所有的只包含中文的用户昵称的计数信息。
func TestForAndRange(t *testing.T) {
	var namesCount map[string]int
	targetsCount := make(map[string]int)
	for k, v := range namesCount {
		matched := true
		for _, r := range k {
			if r < '\u4e00' || r > '\u9fbf' {
				matched = false
				break
			}
		}
		if matched {
			targetsCount[k] = v
		}
	}
}

// [改写]改写上面的例子
//  找到第一个非全中文的用户昵称，从内层循环退出，并从外层循环退出
//
func TestForAndRange02(t *testing.T) {
	var namesCount map[string]int
	targetsCount := make(map[string]int)
	for k, v := range namesCount {
		matched := true
		for _, r := range k {
			if r < '\u4e00' || r > '\u9fbf' {
				matched = false
				break
			}
		}
		if matched {
			targetsCount[k] = v
		} else {
			break
		}
	}
}

// `break` 后面可以放置标签，代表跳转到的位置。






