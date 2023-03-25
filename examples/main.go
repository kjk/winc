package main

import (
	"fmt"

	"github.com/kjk/winc"
)

func newButton(parent winc.Controller, s string, y int) *winc.PushButton {
	btn := winc.NewPushButton(parent)
	btn.SetText(s)
	btn.SetPos(0, y)
	return btn
}

func dispatchSamples() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("Hello World Demo")

	y := 0
	addBtn := func(s string, fn func()) *winc.PushButton {
		btn := newButton(mainWindow, s, y)
		_, dy := btn.Size()
		y += dy + 2
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
