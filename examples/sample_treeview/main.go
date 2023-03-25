package main

import (
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
	{
		btn := newButton(mainWindow, "Run ListView", y)
		_, dy := btn.Size()
		y += dy + 2
		btn.OnClick().Bind(func(e *winc.Event) {
			treeViewExample()
		})
	}
	{
		btn := newButton(mainWindow, "Run TopForm", y)
		_, dy := btn.Size()
		y += dy + 2
		btn.OnClick().Bind(func(e *winc.Event) {
			topForm()
		})
	}
	{
		btn := newButton(mainWindow, "Run ContextMenu", y)
		_, dy := btn.Size()
		y += dy + 2
		btn.OnClick().Bind(func(e *winc.Event) {
			contextMenu()
		})
	}

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
}

func main() {
	dispatchSamples()
	winc.RunMainLoop()
}
