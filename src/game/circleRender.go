package game

import (
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func renderCircle() *canvas.Circle {
	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 8,
	}
	circle.StrokeWidth = 5
	return circle
}

func animateCircle() *canvas.Circle {
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.White
	circle.StrokeWidth = 5
	return circle
}
