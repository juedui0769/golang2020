package test

import (
	"fanyi.baidu/baidu"
	"fanyi.baidu/baidu/user"
	"testing"
)

var userInfo = user.UserInfo{"2015063000000001", "12345678"}

func TestBaiduFanyiInput_GenSign(t *testing.T) {

	input := baidu.BaiduFanyiInput{"apple", "en",
		"zh", "", "1435660288", ""}

	sign := input.GenSign(userInfo)

	t.Log(sign)
	t.Log(input)
}

func TestBaiduFanyiInput_GenQueryStr(t *testing.T) {

	input := baidu.NewBaiduFanyiInput("apple", "en", "zh", userInfo)
	q := input.GenQueryStr()
	t.Log(input)
	t.Log(q)
}

func TestBaiduFanyiInput_GenQueryUrl(t *testing.T) {
	input := baidu.NewBaiduFanyiInput("apple", "en", "zh", userInfo)
	t.Log(input)
	t.Log(input.GenQueryUrl())
}
