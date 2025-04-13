package exercise

import (
	"github.com/rivo/tview"
)

func NewListView(app *tview.Application, title string, border bool, fun1 func(*tview.Application)) *tview.List {
	v := tview.NewList().
		AddItem("List item 1", "", 'a', nil).
		AddItem("List item 2", "", 'b', nil).
		AddItem("List item 3", "", 'c', nil).
		AddItem("List item 4", "", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			fun1(app)
		})
	v.SetBorder(border).SetTitle(title)

	return v
}
