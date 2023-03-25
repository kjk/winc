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
	addBtn := func(s string) *winc.PushButton {
		btn := newButton(mainWindow, s, y)
		_, dy := btn.Size()
		y += dy + 2
		return btn
	}

	{
		btn := addBtn("Run ListView")
		btn.OnClick().Bind(func(e *winc.Event) {
			treeViewExample()
		})
	}
	{
		btn := addBtn("Run TopForm")
		btn.OnClick().Bind(func(e *winc.Event) {
			topForm()
		})
	}
	{
		btn := addBtn("Run ContextMenu")
		btn.OnClick().Bind(func(e *winc.Event) {
			contextMenu()
		})
	}
	{
		btn := addBtn("Run ListView")
		btn.OnClick().Bind(func(e *winc.Event) {
			listView()
		})
	}

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
