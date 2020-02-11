package util

import (
	"crypto/md5"
	"fmt"
)

// 获得`originStr`的MD5值
func GetMd5Str(originStr string) string {
	data := []byte(originStr)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
