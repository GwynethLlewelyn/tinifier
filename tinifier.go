package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

func main() {
	a := app.New()
	wClock := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	wClock.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	wClock.Show()

	w1 := a.NewWindow("Hello World")

	w1.SetContent(widget.NewLabel("Hello World!"))
	w1.Show()

	w2 := a.NewWindow("Larger")
	w2.Resize(fyne.NewSize(100, 100))
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.Resize(fyne.NewSize(120, 120))
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))

	wType := a.NewWindow("Hello Person")

	wType.SetContent(container.NewVBox(makeUI()))
	wType.Show()

	a.Run()
}
