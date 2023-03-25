package main

import (
	"fmt"

	"github.com/kjk/winc"
)

type Item5 struct {
	T string
}

func (item Item5) Text() string    { return item.T }
func (item Item5) ImageIndex() int { return 0 }

func docking() {
	//winc.Init()

	mainWindow := winc.NewForm(nil)
	dock := winc.NewSimpleDock(mainWindow)
	mainWindow.SetLayout(dock)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	fileMn.AddItem("New", winc.NoShortcut)
	editMn := menu.AddSubMenu("Edit")
	cutMn := editMn.AddItem("Cut", winc.Shortcut{winc.ModControl, winc.KeyX})
	copyMn := editMn.AddItem("Copy", winc.NoShortcut)
	pasteMn := editMn.AddItem("Paste", winc.NoShortcut)
	menu.Show()
	copyMn.SetCheckable(true)
	copyMn.SetChecked(true)
	pasteMn.SetEnabled(false)

	tabs := winc.NewMultiPanel(mainWindow)
	tabs.SetPos(10, 10)
	tabs.SetSize(100, 92)

	tree := winc.NewTreeView(mainWindow)
	tree.SetPos(10, 80)
	p := &Item5{"First Item"}
	tree.InsertItem(p, nil, nil)
	sec := &Item5{"Second"}
	if err := tree.InsertItem(sec, p, nil); err != nil {
		panic(err)
	}
	if err := tree.InsertItem(&Item5{"Third"}, p, nil); err != nil {
		panic(err)
	}
	if err := tree.InsertItem(&Item5{"Fourth"}, p, nil); err != nil {
		panic(err)
	}
	for i := 0; i < 50; i++ {
		if err := tree.InsertItem(&Item5{"after second"}, sec, nil); err != nil {
			panic(err)
		}
	}
	tree.Expand(p)
	tree.OnCollapse().Bind(func(e *winc.Event) {
		println("collapse")
	})

	cutMn.OnClick().Bind(func(e *winc.Event) {
		println("cut click")
		ok := tree.EnsureVisible(p)
		fmt.Println("result of EnsureVisible", ok)
	})

	panel := winc.NewPanel(tabs)
	tabs.AddPanel(panel)

	panelDock := winc.NewSimpleDock(panel)
	panel.SetLayout(panelDock)
	panel.SetPos(0, 0)

	panelErr := winc.NewErrorPanel(panel)
	panelErr.SetPos(140, 10)
	panelErr.SetSize(200, 32)
	panelErr.ShowAsError(false)

	edt := winc.NewEdit(panel)
	edt.SetPos(10, 535)
	edt.SetText("some text")

	btn := winc.NewPushButton(panel, "Button")
	// btn.SetSize(100, 40)
	btn.OnClick().Bind(func(e *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})
	btn.SetResIcon(15)

	panelDock.Dock(btn, winc.Top)
	panelDock.Dock(edt, winc.Top)
	panelDock.Dock(panelErr, winc.Top)

	dock.Dock(tree, winc.Left)
	dock.Dock(tabs, winc.Top)

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
}
