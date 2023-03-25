package main

import (
	"fmt"

	"github.com/kjk/winc"
)

func dispatchSamples() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 100)
	mainWindow.SetText("Hello World Demo")

	y := 4
	maxDx := 0
	addBtn := func(s string, fn func()) *winc.PushButton {
		btn := winc.NewPushButton(mainWindow)
		btn.SetText(s)
		btn.SetPos(8, y)
		dx, dy := btn.Size()
		dx = dx * 2
		btn.SetSize(dx, dy)
		y = y + dy + 2
		if dx > maxDx {
			maxDx = dx
		}
		if fn != nil {
			btn.OnClick().Bind(func(e *winc.Event) {
				fn()
			})
		}
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
	dispatchSamples()
	winc.RunMainLoop()
}
