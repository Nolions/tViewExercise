package main

import (
	"fmt"
	"github.com/rivo/tview"
	"tViewExercise/model"
	"tViewExercise/ui"
)

func main() {
	app := tview.NewApplication()

	conf := model.NewAWSConfig()

	pages := tview.NewPages()

	credentials := ui.CredentialsLayout(app, pages, "m", stopApp, switchPage, conf)

	m := tview.NewModal().
		SetText(fmt.Sprintf("exercise")).
		AddButtons([]string{"Next", "Quit"}).
		SetDoneFunc(func(i int, v string) {
			switch i {
			case 0:
				pages.SwitchToPage("listview")
			case 1:
				app.Stop()
			}
		})

	pages.AddPage("credentials", credentials, true, true).
		AddPage("m", m, false, false)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func switchPage(pages *tview.Pages, pageName string) {
	pages.SwitchToPage(pageName)
}

func stopApp(app *tview.Application) {
	app.Stop()
}
