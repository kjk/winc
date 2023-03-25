package main

import (
	"fmt"

	"github.com/kjk/winc"
)

type Item3 struct {
	T       []string
	checked bool
}

func (item Item3) Text() []string    { return item.T }
func (item *Item3) SetText(s string) { item.T[0] = s }

func (item Item3) Checked() bool            { return item.checked }
func (item *Item3) SetChecked(checked bool) { item.checked = checked }
func (item Item3) ImageIndex() int          { return 0 }

func listView() {
	mainWindow := winc.NewForm(nil)
	dock := winc.NewSimpleDock(mainWindow)

	mainWindow.SetSize(700, 600)
	mainWindow.SetText("Controls Demo")

	none := winc.Shortcut{}

	imlist := winc.NewImageList(16, 16)
	imlist.AddResIcon(15)
	imlist.AddResIcon(10)
	imlist.AddResIcon(12)

	ls := winc.NewListView(mainWindow)
	ls.SetImageList(imlist)
	ls.EnableEditLabels(false)
	ls.SetCheckBoxes(true)
	//ls.EnableFullRowSelect(true)
	//ls.EnableHotTrack(true)
	//ls.EnableSortHeader(true)
	//ls.EnableSortAscending(true)

	ls.AddColumn("One", 120)
	ls.AddColumn("Two", 120)
	ls.SetPos(10, 180)
	p1 := &Item3{[]string{"First Item", "A"}, true}
	ls.AddItem(p1)
	p2 := &Item3{[]string{"Second Item", "B"}, true}
	ls.AddItem(p2)
	p3 := &Item3{[]string{"Third Item", "C"}, true}
	ls.AddItem(p3)
	for i := 0; i < 200; i++ {
		p4 := &Item3{[]string{"Fourth Item", "D"}, false}
		ls.AddItem(p4)
	}

	menu := mainWindow.NewMenu()
	fileMn := menu.AddSubMenu("File")
	fileMn.AddItem("New", none)
	editMn := menu.AddSubMenu("Edit")
	delMn := editMn.AddItem("Delete", winc.Shortcut{winc.ModControl, winc.KeyX})
	delAllMn := editMn.AddItem("Delete All", none)
	menu.Show()

	ls.OnEndLabelEdit().Bind(func(e *winc.Event) {
		println("edited", e)
		// acccept label edit event!
		//d := e.Data.(*winc.LabelEditEventData)
		//d.Item.SetText(d.Text)
		//fmt.Println(d.Item.Text())
	})

	delMn.OnClick().Bind(func(e *winc.Event) {
		items := ls.SelectedItems()
		for _, it := range items {
			fmt.Println(it)
		}
	})

	delAllMn.OnClick().Bind(func(e *winc.Event) {
		ls.DeleteAllItems()
	})

	ls.OnClick().Bind(func(e *winc.Event) {
		println("onClick listview")
	})

	dock.Dock(ls, winc.Fill)

	mainWindow.Center()
	mainWindow.Show()
	closeMyself := func(e *winc.Event) {
		mainWindow.Close()
	}
	mainWindow.OnClose().Bind(closeMyself)
}
