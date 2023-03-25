package main

import (
	"fmt"

	"github.com/kjk/winc"
	"github.com/kjk/winc/w32"
)

func push[T any](a *[]T, el T) {
	*a = append(*a, el)
}

func dispatchSamples() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 100)
	mainWindow.SetText("Hello World Demo")

	var buttons []*winc.PushButton
	y := 4
	maxDx := 0
	addBtn := func(s string, fn func()) *winc.PushButton {
		btn := winc.NewPushButton(mainWindow, s)
		btn.SetPos(8, y)
		dx, dy := btn.Size()
		y = y + dy + 2
		if dx > maxDx {
			maxDx = dx
		}
		if fn != nil {
			btn.OnClick().Bind(func(e *winc.Event) {
				fn()
			})
		}
		push(&buttons, btn)
		return btn
	}

	addBtn("Run ListView", treeViewExample)
	addBtn("Run TopForm", topForm)
	addBtn("Run ContextMenu", contextMenu)
	addBtn("Run ListView", listView)
	addBtn("Run Tabs", tabs)
	addBtn("Run Docking", docking)
	addBtn("Run Image", image)
	addBtn("Run ImageBox", imageBox)
	addBtn("Run ScrollView", scrollView)
	addBtn("Run Slider", slider)
	addBtn("Run SplitView", splitView)

	// make all buttons same size
	for _, btn := range buttons {
		_, dy := btn.Size()
		btn.SetSizePx(maxDx, dy)
	}

	// dx, _ := mainWindow.Size()
	mainWindow.SetSize(maxDx+8+8, y-2)

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
}

func main() {
	println("Hello!\n")
	defer func() {
		fmt.Printf("defer!\n")
	}()
	ncm := w32.GetNonClientMetrics()
	fh := ncm.LfMessageFont.Height
	fmt.Printf("Message font height: %v\n", fh)
	dispatchSamples()
	winc.RunMainLoop()
}
