package ui

import (
	"github.com/rivo/tview"
	"tViewExercise/aws"
	"tViewExercise/model"
)

func CredentialsForm(
	app *tview.Application,
	pages *tview.Pages,
	conf *model.AWSConfig,
	switchTo string,
	exitFun func(app *tview.Application),
) *tview.Form {
	form := tview.NewForm()
	form.AddDropDown("Region", aws.Regions, conf.Region, nil).
		AddInputField("AccessKey", conf.AccessKey, 35, nil, nil).
		AddInputField("SecretKey", conf.SecretKey, 35, nil, nil).
		AddInputField("Bucket", conf.Bucket, 35, nil, nil).
		AddCheckbox("Acl", conf.Acl, nil).
		AddButton("Save", func() {
			pages.SwitchToPage(switchTo)
		}).
		AddButton("Reset", func() {
			conf.AccessKey = ""
			conf.SecretKey = ""
			conf.Bucket = ""
			form.GetFormItem(1).(*tview.InputField).SetText("")
			form.GetFormItem(2).(*tview.InputField).SetText("")
			form.GetFormItem(3).(*tview.InputField).SetText("")
		}).
		AddButton("Exit", func() {
			exitFun(app)
		})

	form.SetTitle("Credentials").SetBorder(true)

	return form
}
