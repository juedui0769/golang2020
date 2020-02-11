package main

import (
	"fanyi.baidu/util"
	"testing"
)

func TestCurrentFilePath(t *testing.T) {
	t.Log(util.GetCurrFileDir())
}
