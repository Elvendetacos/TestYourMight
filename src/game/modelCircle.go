package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func paintCircle(ac *canvas.Circle, c *canvas.Circle) {
	for {
		time.Sleep(60 * time.Millisecond)
		currenSize := ac.Size()
		colider := c.Size()
		if currenSize.Width <= colider.Width && currenSize.Height <= colider.Height {
			ac.FillColor = color.Gray{Y: 100}
			ac.StrokeColor = color.Gray{Y: 100}
		} else {
			ac.FillColor = color.Gray{Y: 0x99}
			ac.StrokeColor = color.Gray{Y: 0x99}
		}
	}
}

func growCircle(ac *canvas.Circle) {
	for {
		time.Sleep(60 * time.Millisecond)
		currentSize := ac.Size()
		if currentSize.Width <= 500 {
			newWidth := currentSize.Width + 10
			newHeight := currentSize.Height + 10
			center := ac.Position().Add(fyne.NewPos(currentSize.Width/2, currentSize.Height/2))
			ac.Resize(fyne.NewSize(newWidth, newHeight))
			ac.Move(center.Subtract(fyne.NewPos(newWidth/2, newHeight/2)))
		}
	}
}

func progressbar(setTime *widget.ProgressBar, ac *canvas.Circle, c *canvas.Circle, mainMenu fyne.Window, content *fyne.Container) {
	timer := 1.0
	for timers := timer; timers >= 0.0; timers -= 0.1 {
		time.Sleep(1 * time.Second)
		setTime.SetValue(timers)
	}
	currentSize := ac.Size()
	colider := c.Size()
	if currentSize.Width <= colider.Width && currentSize.Height <= colider.Height {
		screen(mainMenu, content, "Ganaste :3")
	} else {
		screen(mainMenu, content, "Perdiste :(")
	}
}

func shrinkCircle(currentSize fyne.Size, ac *canvas.Circle) {
	if currentSize.Width > 15 && currentSize.Height > 15 {
		newWidth := currentSize.Width - 15
		newHeight := currentSize.Height - 15
		center := ac.Position().Add(fyne.NewPos(currentSize.Width/2, currentSize.Height/2))
		ac.Resize(fyne.NewSize(newWidth, newHeight))
		ac.Move(center.Subtract(fyne.NewPos(newWidth/2, newHeight/2)))
	}
}
