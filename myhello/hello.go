package main

import (
	"fanyi.baidu/util"
	"fmt"
	"myhello/input"
)

func main() {
	demo02()
}

func demo02() {
	input.InputDemo()
}

func demo01() {
	fmt.Printf("hello, world! welcome to golang!\n")

	c := 5 + 5i
	fmt.Printf("Value is : %v\n", c)

	fmt.Println(util.GetCurrFileDir())

	fmt.Scanf("%s")

	fmt.Println("结束...")
}