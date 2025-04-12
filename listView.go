package main

import "github.com/rivo/tview"

func newListView(app *tview.Application, title string, border bool) *tview.List {
	v := tview.NewList().
		AddItem("List item 1", "", 'a', echo("a")).
		AddItem("List item 2", "", 'b', echo("b")).
		AddItem("List item 3", "", 'c', echo("c")).
		AddItem("List item 4", "", 'd', echo("b")).
		AddItem("Quit", "Press to exit", 'q', stopApp(app))
	v.SetBorder(border).SetTitle(title)

	return v
}
