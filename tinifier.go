package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
//	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// UI elements

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

//
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

	wMain := a.NewWindow(os.Args[0])
	wMain.Resize(fyne.NewSize(800.0, 600.0))

	image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	// image := canvas.NewImageFromFile(fileName)
	image.SetMinSize(fyne.NewSize(800.0, 600.0))
	image.FillMode = canvas.ImageFillOriginal
	wMain.SetContent(image)

	// Open a dialog box for getting a file.
	onChosen := func(f fyne.URIReadCloser, err error) {
		if err != nil {
			fyne.LogError("file dialog", err)
			return
		}
		if f == nil {
			fyne.LogError("file dialog", fmt.Errorf("unknown or unexisting file"))
			return
		}
		fmt.Printf("chosen: %v\n", f.URI())
		// loadFile(wMain, f, fileLength(f.URI().Path()))
		image = canvas.NewImageFromURI(f.URI()) // TODO work with library upstream to not do this
		image.FillMode = canvas.ImageFillOriginal
		wMain.SetContent(image)
		wMain.Show()
	}

	d := dialog.NewFileOpen(onChosen, wMain)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".jpg", ".jpeg", ".webp", ".png", ".avif"}))
	// d.Resize(fyne.NewSize(300.0, 200.0))
	d.Show()

	wMain.Show()

	a.Run()
}

// additional functions

// Returns the file length as an int64, or zero if file not found.
func fileLength(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0
	}

	return info.Size()
}

func loadFile(parent fyne.Window, r io.ReadCloser, length int64) {
	var data []byte
	bytes, err := r.Read(data)
	if err != nil {
		dialog.ShowError(err, parent)
		return
	}

	_ = r.Close()

	if bytes < int(length) {
		dialog.ShowError(fmt.Errorf("%d out of expected %d byte(s) read! File may be corrupt", bytes, length), parent)
		return
	}
	dialog.ShowInformation("Status", fmt.Sprintf("%d byte(s) read", bytes), parent)
}