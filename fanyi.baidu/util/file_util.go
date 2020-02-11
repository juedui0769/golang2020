package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const FANYI_BAIDU_CONFIG_FILE string = "fanyi.baidu.conf"

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

// https://www.jianshu.com/p/93772ececf65
// 判断文件是否存在
// GetCurrFileDir() + "/" + FANYI_BAIDU_CONFIG_FILE
func IfFileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		//log.Logger{}(err)
		//log.Println(err)
		return false
	}
	return true
}

// 按行读取文件
func ReadLines(filename string) (lines []string) {

	if IfFileExist(filename) {
		byteContent, err := ioutil.ReadFile(filename)
		if err == nil {
			//t.Log(string(byteContent))
			strContent  := string(byteContent)
			strArr := strings.Split(strContent, "\n")
			for _, str := range strArr {
				//t.Logf("%d : %s", index, strings.TrimSpace(str))
				if strings.TrimSpace(str) != "" {
					lines = append(lines, str)
				}
			}
		}
	}

	return lines
}

// 写入文件
// 传入文件名（完整路径），和要写入的内容（切片类型）
func WriteLines(filename string, lines []string) {
	if lines == nil {
		log.Fatal("请传入要写入的文件内容")
		return
	}
	if !IfFileExist(filename) {
		file, err := os.Create(filename)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range lines {
			file.WriteString(line)
			file.WriteString("\r\n")
		}
	}
}

