package test

import (
	"fanyi.baidu/util"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
)

var filePath string = ""

func TestCurrentDirectoryPath(t *testing.T) {
	t.Log(os.Args[0])
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		t.Error(err)
	}
	t.Log(strings.Replace(dir, "\\", "/", -1))
t.Log("===")
	t.Log(util.GetCurrFileDir())
}

// https://studygolang.com/articles/2772
func TestSystemPath(t *testing.T) {
	curr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	t.Log(curr.HomeDir)
	t.Log(curr)
}

// 文件是否存在
// https://www.jianshu.com/p/93772ececf65
func TestFileExist(t *testing.T) {
	info, err := os.Stat(util.GetCurrFileDir()+"/fanyi.baidu.conf")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(info)
}

func TestFileExist01(t *testing.T) {
	file := util.GetCurrFileDir() + "/" + util.FANYI_BAIDU_CONFIG_FILE
	flag := util.IfFileExist(file)
	t.Log(flag)
}

// 读取文件
func TestReadFile(t *testing.T) {
	file := util.GetCurrFileDir() + "/" + util.FANYI_BAIDU_CONFIG_FILE
	lines := util.ReadLines(file)
	if lines != nil {
		for index, line := range lines {
			t.Logf("%d: %s\n", index, line)
		}
	} else {
		t.Log("lines is null")
	}
}

// 参考，谢孟军，7.5文件操作
func TestWriteFile(t *testing.T){
	filename := util.GetCurrFileDir() + "/" + util.FANYI_BAIDU_CONFIG_FILE
	t.Log(filename)
	lines := []string{
		"app_id:     2020022400027074111",
		"secret_key: ah9YpW6KBz0gXJJOfeee",
	}
	util.WriteLines(filename, lines)
}
