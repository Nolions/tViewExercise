package ui

import (
	"github.com/rivo/tview"
)

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
