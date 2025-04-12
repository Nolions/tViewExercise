package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	lLay := newListView(app, "menu", true)
	//rLay := newTreeView(".")
	//form := newForm("user info", model.NewUser(), true)
	grid := newGrid()

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(lLay, 30, 20, false).
		//AddItem(form, 40, 20, true).
		//AddItem(rLay, 30, 80, false).
		AddItem(grid, 90, 80, false)

	frame := netFrame("frame Demo", true, flex)

	if err := app.SetRoot(frame, true).Run(); err != nil {
		panic(err)
	}
}

func echo(text string) func() {
	return func() {
		//fmt.Println(text)
	}
}

func stopApp(app *tview.Application) func() {
	return func() {
		app.Stop()
	}
}

func newBox(title string, border bool) *tview.Box {
	v := tview.NewBox().
		SetBorder(border).
		SetTitle(title)

	return v
}
