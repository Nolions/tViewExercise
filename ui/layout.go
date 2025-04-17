package ui

import (
	"github.com/rivo/tview"
	"tViewExercise/model"
)

type AppContext struct {
	App     *tview.Application
	Pages   *tview.Pages
	AwsConf *model.AWSConfig
}

func NewAppContext(app *tview.Application, pages *tview.Pages, conf *model.AWSConfig) *AppContext {
	return &AppContext{
		App:     app,
		Pages:   pages,
		AwsConf: conf,
	}
}

// SetFocusOnPage
// 泛用 Focus 切換器
func SetFocusOnPage(app *tview.Application, pageName string, focusMap map[string]tview.Primitive) {
	if view, ok := focusMap[pageName]; ok && view != nil {
		app.SetFocus(view)
	}
}

// WrapCentered
// 通用：將元件置中包住（使用 Flex）
func WrapCentered(content tview.Primitive) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(nil, 0, 1, false).
				AddItem(content, 50, 1, true).
				AddItem(nil, 0, 1, false),
			0, 2, true).
		AddItem(nil, 0, 1, false)
}
