package test

import (
	"fanyi.baidu/baidu"
	"fanyi.baidu/baidu/user"
	"math/rand"
	"testing"
)

func TestBaiduFanyiInput_GenSign(t *testing.T) {
	userInfo := user.UserInfo{"123", "456"}
	input := baidu.BaiduFanyiInput{"apple", "en",
		"zh", userInfo.AppId, string(rand.Intn(100000)), ""}

	sign := input.GenSign(userInfo)
	input.Sign = sign
	t.Log(sign)
}
