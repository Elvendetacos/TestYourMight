package ui

import (
	"SpeedPlataform/src/game"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func RenderMenu(mainMenu fyne.Window) {
	mainMenu.Resize(fyne.NewSize(700, 600))
	menuContent := container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 80)))
	content := container.New(layout.NewCenterLayout(), menuContent)
	StartGame := widget.NewButtonWithIcon("Iniciar el juego", theme.ComputerIcon(), func() {
		game.RenderGame(mainMenu, content)
	})

	options := widget.NewButtonWithIcon("Opciones", theme.SettingsIcon(), func() {
		renderInstructions(mainMenu, content)
	})

	exitButton := widget.NewButtonWithIcon("Salir", theme.ContentClearIcon(), func() {
		mainMenu.Close()
	})
	menuContent.Add(StartGame)
	menuContent.Add(options)
	menuContent.Add(exitButton)
	mainMenu.SetContent(content)
}

func renderInstructions(mainMenu fyne.Window, content *fyne.Container) {
	text := canvas.NewText("Aquí debería de haber algo, pero las opciones son mentales", color.White)
	text.TextStyle = fyne.TextStyle{Bold: true}
	goBack := widget.NewButton("Regresar", func() {
		mainMenu.SetContent(content)
	})
	menuOptions := container.New(layout.NewGridWrapLayout(fyne.NewSize(400, 80)), text, goBack)
	centerMenu := container.New(layout.NewCenterLayout(), menuOptions)
	mainMenu.SetContent(centerMenu)
}
