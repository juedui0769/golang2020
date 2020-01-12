package ch04

import (
	"fmt"
	"testing"
)

func TestDefer601(t *testing.T) {
	func (){
		for i:=0;i<5;i++{
			defer func(){
				fmt.Printf("%d ", i)
			}()
		}
	}()
}

// 上面输出全部是 5 ，
// 下面 4 3 2 1 0

func TestDefer602(t *testing.T) {
	func (){
		for i:=0;i<5;i++{
			defer func(i int){
				fmt.Printf("%d ", i)
			}(i)
		}
	}()
}

