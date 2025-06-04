//go:generate fyne bundle -o bundled.go assets/QuickSort-logo-small.png
//go:generate fyne bundle -o bundled.go -append assets/QuickSort-logo.png
//go:generate fyne bundle -o bundled.go -append assets/style.css
package main

import (
	"fmt"
	"math/rand"
	//	"sort"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const dataSize = 50
const goroutineThreshold = 5

var (
	data  []int
	rects []*canvas.Rectangle
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func main() {
	a := app.New()
	w := a.NewWindow("Goroutine QuickSort Visualisation")
	w.Resize(fyne.NewSize(800.0, 400.0))

	image := canvas.NewImageFromResource(resourceQuickSortLogoSmallPng)

	image.SetMinSize(fyne.NewSize(800.0, 600.0))
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	// rand.Seed(time.Now().UnixNano()) // deprecated
	data = rand.Perm(dataSize)

	// Create the rectangles to represent data
	rects = make([]*canvas.Rectangle, dataSize)
	bars := make([]fyne.CanvasObject, dataSize)
	for i := range data {
		rect := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
		height := float32(data[i]+1) * 5
		rect.SetMinSize(fyne.NewSize(10, height))
		rects[i] = rect
		bars[i] = rect
	}

	barContainer := container.NewHBox(bars...)
	btn := widget.NewButton("Sort!", func() {
		go func() {
			QuickSort(0, len(data)-1)
			wg.Wait()
			fmt.Println("Sorting done!")
		}()
	})

	layout := container.NewBorder(nil, btn, nil, nil, barContainer)
	w.SetContent(layout)
	w.ShowAndRun()
}

func QuickSort(low, high int) {
	if low < high {
		pivot := partition(low, high)

		if (high - low) > goroutineThreshold {
			wg.Add(2)
			go func() {
				defer wg.Done()
				QuickSort(low, pivot-1)
			}()
			go func() {
				defer wg.Done()
				QuickSort(pivot+1, high)
			}()
		} else {
			QuickSort(low, pivot-1)
			QuickSort(pivot+1, high)
		}
	}
}

func partition(low, high int) int {
	pivot := data[high]
	i := low - 1

	for j := low; j < high; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
			updateUI(i, j)
		}
	}

	data[i+1], data[high] = data[high], data[i+1]
	updateUI(i+1, high)
	return i + 1
}

func updateUI(i, j int) {
	mutex.Lock()
	defer mutex.Unlock()

	fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Swap", Content: fmt.Sprintf("Swapped %d and %d", i, j)})

	rects[i].FillColor = theme.Color(theme.ColorNameWarning)
	rects[j].FillColor = theme.Color(theme.ColorNameError)
	rects[i].Refresh()
	rects[j].Refresh()

	rects[i].SetMinSize(fyne.NewSize(10, float32(data[i]+1)*5))
	rects[j].SetMinSize(fyne.NewSize(10, float32(data[j]+1)*5))

	time.Sleep(50 * time.Millisecond)

	rects[i].FillColor = theme.Color(theme.ColorNamePrimary)
	rects[j].FillColor = theme.Color(theme.ColorNamePrimary)
	rects[i].Refresh()
	rects[j].Refresh()
}
