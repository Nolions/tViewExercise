package main

import (
	"github.com/rivo/tview"
	"tViewExercise/model"
	"tViewExercise/ui"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	conf := model.NewAWSConfig()

	// credentials 頁面
	credentialsForm := ui.CredentialsForm(app, pages, conf, "manager", func(app *tview.Application) {
		app.Stop()
	})
	credentialsPage := ui.WrapCentered(credentialsForm)

	// manager 頁面
	managerPage := ui.ManagerLayout(app, pages)
	browserLayout := managerPage.GetItem(1).(*tview.Flex) // 第二個是 browserLayout

	pages.AddPage("credentials", credentialsPage, true, true)
	pages.AddPage("manager", managerPage, true, false)

	focusMap := map[string]tview.Primitive{
		"credentials": credentialsForm.GetFormItem(1).(tview.Primitive), // AccessKey input
		"manager":     browserLayout,
	}

	ui.SetFocusOnPage(app, "credentials", focusMap)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
