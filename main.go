package main

import (
	"fmt"
	"github.com/rivo/tview"
	"tViewExercise/widget"
)

func main() {
	app := tview.NewApplication()
	lLay := widget.NewListView(app, "menu", true, stopApp)
	//rLay := widget.NewTreeView(".")
	//form := widget.NewForm("user info", model.NewUser(), true)
	//grid := widget.NewGrid()

	//flex := tview.NewFlex().
	//	SetDirection(tview.FlexColumn).
	//	AddItem(lLay, 30, 20, false).
	//	//AddItem(form, 40, 20, true).
	//	//AddItem(rLay, 30, 80, false).
	//	AddItem(grid, 90, 80, false)
	//
	//frame := netFrame("frame Demo", true, flex)

	//pages := newPages(app, []string{"Next", "Quit"}, 3, quickAction, changePageAction)

	pages := tview.NewPages()

	m := tview.NewModal().
		SetText(fmt.Sprintf("demo")).
		AddButtons([]string{"Next", "Quit"}).
		SetDoneFunc(func(i int, v string) {
			switch i {
			case 0:
				pages.SwitchToPage("listview")
			case 1:
				app.Stop()
			}
		})

	pages.AddPage("model", m, false, true).
		AddPage("listview", lLay, false, false)

	if err := app.SetRoot(pages, false).Run(); err != nil {
		panic(err)
	}
}

func changePageAction(pages *tview.Pages, page, pageCount int) {
	nextPage := fmt.Sprintf("page-%d", (page+1)%pageCount)
	pages.SwitchToPage(nextPage)
}

func quickAction(app *tview.Application) {
	app.Stop()
}

func stopApp(app *tview.Application) {
	app.Stop()
}

func newBox(title string, border bool) *tview.Box {
	v := tview.NewBox().
		SetBorder(border).
		SetTitle(title)

	return v
}
