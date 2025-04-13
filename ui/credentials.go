package ui

import (
	"github.com/rivo/tview"
	"tViewExercise/aws"
	"tViewExercise/model"
)

func CredentialsLayout(
	app *tview.Application,
	pages *tview.Pages,
	pageName string,
	stopFun func(*tview.Application),
	switchFun func(pages *tview.Pages, pageName string),
	conf *model.AWSConfig,
) *tview.Flex {
	form := CredentialsForm(app, pages, conf, pageName, stopFun, switchFun)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false). // 上方空白
		AddItem(
			tview.NewFlex().
				AddItem(nil, 0, 1, false).  // 左側空白
				AddItem(form, 50, 1, true). // 中央表單寬度50
				AddItem(nil, 0, 1, false),  // 右側空白
						0, 2, true). // 中間高度自動分配
		AddItem(nil, 0, 1, false) // 下方空白

	return layout
}

func CredentialsForm(
	app *tview.Application,
	pages *tview.Pages,
	conf *model.AWSConfig,
	pageName string,
	exitFun func(*tview.Application),
	switchFun func(pages *tview.Pages, pageName string),
) *tview.Form {
	form := tview.NewForm().
		AddDropDown("Region", aws.Regions, conf.Region, nil).
		AddInputField("AccessKey", conf.AccessKey, 35, nil, nil).
		AddInputField("SecretKey", conf.SecretKey, 35, nil, nil).
		AddInputField("Bucket", conf.Bucket, 35, nil, nil).
		AddCheckbox("Acl", conf.Acl, nil).
		AddButton("Save", func() {
			switchFun(pages, pageName)
		}).
		AddButton("Reset", func() {

		}).
		AddButton("Exit", func() {
			exitFun(app)
		})

	form.SetTitle("Credentials").SetBorder(true)

	return form
}
