package template

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"time"
)

func Screen(mainMenu fyne.Window, content *fyne.Container, text string) {
	textFail := canvas.NewText(text, color.White)
	textFail.TextSize = 32
	centerText := container.New(layout.NewCenterLayout(), textFail)
	mainMenu.SetContent(centerText)
	time.Sleep(1 * time.Second)
	mainMenu.SetContent(content)
}
