package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	tinifierApp := app.New()
	w := tinifierApp.NewWindow("Tinifier")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
			hello,
			widget.NewButton("Hi!", func() {
					hello.SetText("Welcome 😀")
			}),
	))

	w.ShowAndRun()
}