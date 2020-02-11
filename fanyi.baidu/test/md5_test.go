package test

import (
	"crypto/md5"
	"fanyi.baidu/util"
	"fmt"
	"io"
	"testing"
)

var inputStr = "2015063000000001apple143566028812345678"

func TestMd5V2(t *testing.T) {
	w := md5.New()
	io.WriteString(w, inputStr)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	t.Log(md5Str)
}

func TestGetMd5Str(t *testing.T) {
	t.Log(util.GetMd5Str(inputStr))
}

func TestMd5(t *testing.T) {
	data := []byte(inputStr)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	t.Log(md5Str)
}
