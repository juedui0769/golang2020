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

func TestFileExist(t *testing.T) {

}




