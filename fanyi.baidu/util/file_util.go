package util

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// 获得当前目录
func GetCurrFileDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获得用户目录
// https://studygolang.com/articles/2772
func GetUserHomeDir() string {
	curr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return curr.HomeDir
}

