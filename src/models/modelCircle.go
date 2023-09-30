package models

import (
	"TestYourMight/src/template"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

type GameCircle struct {
	ac       *canvas.Circle
	c        *canvas.Circle
	setTime  *widget.ProgressBar
	mainMenu fyne.Window
	content  *fyne.Container
}

func NewGameCircle(ac, c *canvas.Circle, setTime *widget.ProgressBar, mainMenu fyne.Window, content *fyne.Container) *GameCircle {
	return &GameCircle{
		ac:       ac,
		c:        c,
		setTime:  setTime,
		mainMenu: mainMenu,
		content:  content,
	}
}

func (gc *GameCircle) PaintCircle() {
	for {
		time.Sleep(60 * time.Millisecond)
		currentSize := gc.ac.Size()
		collider := gc.c.Size()
		if currentSize.Width <= collider.Width && currentSize.Height <= collider.Height {
			gc.ac.FillColor = color.Gray{Y: 100}
			gc.ac.StrokeColor = color.Gray{Y: 100}
		} else {
			gc.ac.FillColor = color.Gray{Y: 0x99}
			gc.ac.StrokeColor = color.Gray{Y: 0x99}
		}
	}
}

func (gc *GameCircle) GrowCircle() {
	for {
		time.Sleep(60 * time.Millisecond)
		currentSize := gc.ac.Size()
		if currentSize.Width <= 500 {
			newWidth := currentSize.Width + 10
			newHeight := currentSize.Height + 10
			center := gc.ac.Position().Add(fyne.NewPos(currentSize.Width/2, currentSize.Height/2))
			gc.ac.Resize(fyne.NewSize(newWidth, newHeight))
			gc.ac.Move(center.Subtract(fyne.NewPos(newWidth/2, newHeight/2)))
		}
	}
}

func (gc *GameCircle) Progressbar() {
	timer := 1.0
	for timers := timer; timers >= 0.0; timers -= 0.1 {
		time.Sleep(1 * time.Second)
		gc.setTime.SetValue(timers)
	}
	currentSize := gc.ac.Size()
	collider := gc.c.Size()
	if currentSize.Width <= collider.Width && currentSize.Height <= collider.Height {
		template.Screen(gc.mainMenu, gc.content, "Ganaste :3")
	} else {
		template.Screen(gc.mainMenu, gc.content, "Perdiste :(")
	}
}

func (gc *GameCircle) ShrinkCircle(currentSize fyne.Size) {
	if currentSize.Width > 15 && currentSize.Height > 15 {
		newWidth := currentSize.Width - 15
		newHeight := currentSize.Height - 15
		center := gc.ac.Position().Add(fyne.NewPos(currentSize.Width/2, currentSize.Height/2))
		gc.ac.Resize(fyne.NewSize(newWidth, newHeight))
		gc.ac.Move(center.Subtract(fyne.NewPos(newWidth/2, newHeight/2)))
	}
}
