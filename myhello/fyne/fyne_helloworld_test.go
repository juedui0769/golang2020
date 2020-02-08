package fyne

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"testing"
)

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
