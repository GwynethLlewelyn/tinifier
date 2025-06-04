//go:generate fyne bundle -o bundled.go assets/tinifier-logo.png
//go:generate fyne bundle -o bundled.go -append assets/style.css
package main

import (
	"fmt"
	//	"io"
	"os"
	//	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	//	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	wMain := a.NewWindow(os.Args[0])
	wMain.Resize(fyne.NewSize(800.0, 600.0))

	image := canvas.NewImageFromResource(resourceTinifierLogoPng)

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
		image = canvas.NewImageFromURI(f.URI()) // TODO work with library upstream to not do this
		image.FillMode = canvas.ImageFillOriginal
		wMain.SetContent(image)
		wMain.Show()
	}

	d := dialog.NewFileOpen(onChosen, wMain)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".jpg", ".jpeg", ".webp", ".png", ".avif"}))
	d.Show()

	wMain.Show()

	a.Run()
}
