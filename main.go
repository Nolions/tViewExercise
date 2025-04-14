package main

import (
	"github.com/rivo/tview"
	"tViewExercise/model"
	"tViewExercise/ui"
)

func main() {
	app := tview.NewApplication()

	conf := model.NewAWSConfig()

	pages := tview.NewPages()

	credentials := ui.CredentialsLayout(app, pages, "manager", stopApp, switchPage, conf)
	alert := ui.AlertModel("alert", "test/n test", pages, "credentials", switchPage)
	m := ui.ManagerLayout(app)

	pages.AddPage("credentials", credentials, true, true).
		AddPage("alert", alert, false, false).
		AddPage("manager", m, true, false)

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
