package hello

import (
	"fmt"
	"happygo.wxg.com/study/hellolib"
)

func main()  {
	fmt.Print("你好，世界！\n")
	fmt.Printf("2 和 3中最大的是 %d!\n", hellolib.Max(2, 3))
}