package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func RenderGame(mainMenu fyne.Window, content *fyne.Container) {
	setTime := widget.NewProgressBar()
	c := renderCircle()
	ac := animateCircle()
	mainMenu.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeySpace {
			currentSize := ac.Size()
			shrinkCircle(currentSize, ac)
		}
	})
	trowButton := canvas.NewText("Press spacee", color.White)
	go progressbar(setTime, ac, c, mainMenu, content)
	go growCircle(ac)
	go paintCircle(ac, c)
	drawButton := container.NewVBox(trowButton)
	drawButton.Move(fyne.NewPos(300, 0))
	drawObjects := container.NewCenter(container.NewGridWrap(fyne.NewSize(200, 200), c))
	drawObjects.Move(fyne.NewPos(350, 300))
	drawBar := container.NewVBox(setTime)
	drawBar.Move(fyne.NewPos(0, 555))
	drawBar.Resize(fyne.NewSize(690, 50))
	drawCircle := container.NewCenter(container.NewGridWrap(fyne.NewSize(200, 200), ac))
	drawCircle.Move(fyne.NewPos(350, 300))
	containerRender := container.NewWithoutLayout(drawCircle, drawObjects, drawButton, drawBar)
	mainMenu.SetContent(containerRender)
}
