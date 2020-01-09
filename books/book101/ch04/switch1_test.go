package ch04

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	content := "Python"
	switch content {
	default:
		fmt.Println("Unknown language")
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go":
		fmt.Println("A compiled language")
	}
}

func Test02(t *testing.T) {
	//content := "Python"
	switch content := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go":
		fmt.Println("A compiled language")
	}
}

//func getContent() interface{} {
//	return "Go"
//}

// 上面的返回`interface{}`类型也是可以的。

func getContent() string {
	return "Go"
}

func Test03(t *testing.T) {
	switch content := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Python", "Ruby":
		fmt.Println("A interpreted Language")
	case "Go", "C", "Java":
		fmt.Println("A compiled language")
	}
}

// fallthrough

func Test04(t *testing.T) {
	switch content := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Ruby":
		fallthrough
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go", "C", "Java":
		fmt.Println("A compiled language")
	}
}

// `case`中可以使用`break`

func Test05(t *testing.T) {
	switch content := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Ruby":
		break
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go", "C", "Java":
		fmt.Println("A compiled language")
	}
}






