package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// golang 从键盘读取输入
// (1) https://blog.csdn.net/wangxinwen/article/details/85040126
// (2) https://blog.csdn.net/weixin_34119545/article/details/86134706
// (3*) https://www.jianshu.com/p/bd37b77f0fcf
// (4) https://github.com/eiannone/keyboard
func InputDemo() {
	list := make([]string, 2)
	str := ReadStdin("请输入一个字符串")
	list = append(list, "'" + str + "'")
	str = ReadStdin("请输入一个字符串")
	list = append(list, "'" + str + "'")
	fmt.Println(list)
}

var reader = bufio.NewReader(os.Stdin)

func ReadStdin(tip string) string {
	fmt.Printf("%s: ", tip)
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}