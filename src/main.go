package main

import (
	"SpeedPlataform/src/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	r, _ := fyne.LoadResourceFromPath("assets/Icons/IconGame.png")
	myApp.SetIcon(r)
	menu := myApp.NewWindow("Test your might")
	menu.SetFixedSize(true)
	ui.RenderMenu(menu)
	menu.Show()
	myApp.Run()
}
