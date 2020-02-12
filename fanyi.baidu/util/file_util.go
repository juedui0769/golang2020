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
// (1) 文件不存在时，创建
// (2*) 文件存在时，覆盖更新
func WriteLines(filename string, lines []string) {
	if lines == nil {
		log.Fatal("请传入要写入的文件内容")
		return
	}
	// 如果文件不存在，创建文件，写入内容
	if !IfFileExist(filename) {
		file, err := os.Create(filename)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		writeLines(file, lines)
	} else {
		// 如果文件存在，应该更新文件内容
		// 0777表示：创建了一个普通文件，所有人拥有所有的读、写、执行权限
		// 0666表示：创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
		// 0644表示：创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限
		file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		writeLines(file, lines)
	}
}

// 私有方法
func writeLines(file *os.File, lines []string) {
	for _, line := range lines {
		file.WriteString(line)
		file.WriteString("\r\n")
	}
}




