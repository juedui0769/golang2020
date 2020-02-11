package fyne

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"testing"
)

// 2020年2月11日: 之前测试是可以的, 今天再次测试, 发现有问题，一直在加载，我不得不强制停止。
func TestHello(t *testing.T) {
	t.Log("Hello fyne ...")

	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		})))

	w.ShowAndRun()
}
