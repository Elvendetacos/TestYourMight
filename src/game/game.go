package game

import (
	"TestYourMight/src/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Game struct {
	mainMenu   fyne.Window
	content    *fyne.Container
	ac         *canvas.Circle
	c          *canvas.Circle
	setTime    *widget.ProgressBar
	gameCircle *models.GameCircle
}

func NewGame(mainMenu fyne.Window, content *fyne.Container) *Game {
	ac := animateCircle()
	c := renderCircle()
	setTime := widget.NewProgressBar()
	gameCircle := models.NewGameCircle(ac, c, setTime, mainMenu, content)

	return &Game{
		mainMenu:   mainMenu,
		content:    content,
		ac:         ac,
		c:          c,
		setTime:    setTime,
		gameCircle: gameCircle,
	}
}

func (g *Game) Render() {
	setTime := g.setTime
	ac := g.ac
	c := g.c

	mainMenu := g.mainMenu

	mainMenu.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeySpace {
			currentSize := ac.Size()
			g.gameCircle.ShrinkCircle(currentSize)
		}
	})
	trowButton := canvas.NewText("Press spacee", color.White)
	go g.gameCircle.Progressbar()
	go g.gameCircle.GrowCircle()
	go g.gameCircle.PaintCircle()
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
