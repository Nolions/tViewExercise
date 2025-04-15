package main

import (
	"fmt"
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

	modal := ui.FilePickerModal(filePicker, 60, 20, func() {
		pages.HidePage("filepicker")
	})

	pages.AddPage("credentials", credentialsPage, true, true)
	pages.AddPage("manager", managerPage, true, false)
	pages.AddPage("filepicker", modal, true, false)

	focusMap := map[string]tview.Primitive{
		"credentials": credentialsForm.GetFormItem(1).(tview.Primitive), // AccessKey input
		"manager":     browserLayout,
	}

	ui.SetFocusOnPage(app, "credentials", focusMap)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
