package ch04

import (
	"fmt"
	"testing"
)

func Test201(t *testing.T) {
	//invalid type assertion: int(123).(int) (non-interface type int on left)
	//t.Log(int(123).(int))

	t.Log(interface{}(123).(int))

	// 所以`CheckType`必须接受一个`interface{}`类型的参数
	// `interface{}` 是一个特殊的接口类型，代表空接口。
	// 所有类型都是它的实现类型。
}

func CheckType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("The string is '%s'.\n", v.(string))
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		fmt.Printf("The integer is %d.\n", v)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", v)
	}
}

func Test202(t *testing.T) {
	CheckType2(1)
	CheckType2("hello world")
	CheckType2(true)
	// `byte`, `rune` 也是数值类型
	var x byte = 1
	var y rune = '中'
	CheckType2(x)
	CheckType2(y)
	// `nil` 也可以传入到 `CheckType` 函数中
	CheckType2(nil)
}

// `CheckType2` 是 `CheckType` 更简洁的表达，不用在内部显示的转换。
func CheckType2(v interface{}) {
	switch i := v.(type) {
	case string:
		fmt.Printf("The string is '%s'.\n", i)
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		fmt.Printf("The integer is %d.\n", i)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", i)
	}
}
