package main

import (
	"fmt"
	"github.com/rivo/tview"
	"tViewExercise/model"
	"tViewExercise/ui"
)

func main() {
	app := initApp()
	pages := tview.NewPages()
	conf := model.NewAWSConfig()
	appCtx := ui.NewAppContext(app, pages, conf)

	credentialsPage := appCtx.CredentialsLayout() // credentials 頁面
	managerPage := appCtx.ManagerLayout()         // manager 頁面

	filePicker := ui.FilePickerLayout(ui.FilePickerOption{
		StartDir:          ".",
		AllowFolderSelect: false,
		ExtensionFilter:   []string{},
		OnSelect: func(path string) {
			fmt.Println("你選擇了：", path)
			pages.HidePage("filepicker")
		},
	})
	filePicker.SetBorder(true).SetTitle("Select a file")

	modal := ui.FilePickerModal(filePicker, 60, 15, func() {
		pages.HidePage("filepicker")
	})

	pages.AddPage("credentials", credentialsPage, true, true)
	pages.AddPage("manager", managerPage, true, false)
	pages.AddPage("filepicker", modal, true, false)

	focusMap := map[string]tview.Primitive{
		"credentials": credentialsPage.GetItem(1).(tview.Primitive),
		"manager":     managerPage.GetItem(1).(*tview.Flex),
	}

	ui.SetFocusOnPage(app, "credentials", focusMap)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func initApp() *tview.Application {
	app := tview.NewApplication()
	app.EnableMouse(true)

	return app
}
