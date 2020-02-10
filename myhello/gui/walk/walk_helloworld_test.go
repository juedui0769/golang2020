package walk_test

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"testing"
)

// 运行失败
// D:\golang\lib\src\github.com\akavel\rsrc\rsrc.exe -manifest main.manifest -o main.syso
func TestWalkHelloworld(t *testing.T){
	window, _ := walk.NewMainWindow()

	window.SetTitle("你好世界！")

	window.SetWidth(400)
	window.SetHeight(400)

	// 窗体横坐标 = ( 屏幕宽度 - 窗体宽度 ) / 2
	originX := ( int(win.GetSystemMetrics(0)) - window.Width() ) / 2

	// 窗体纵坐标 = ( 屏幕高度 - 窗体高度 ) / 2
	originY := ( int(win.GetSystemMetrics(1)) - window.Height() ) / 2

	window.SetX(originX)
	window.SetY(originY)

	window.Show()

	window.Run()
}
