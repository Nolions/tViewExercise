package widget

import (
	"fmt"
	"github.com/rivo/tview"
)

func NewModel(
	app *tview.Application,
	labs []string,
	page int,
	pageCount int,
	pages *tview.Pages,
	qAction func(*tview.Application),
	changePage func(pages *tview.Pages, page, pageCount int),
) *tview.Modal {
	m := tview.NewModal().
		SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1))

	if len(labs) > 0 {
		m.AddButtons(labs).
			SetDoneFunc(func(i int, v string) {
				switch i {
				case 0:
					changePage(pages, page, pageCount)
				case 1:
					qAction(app)
				}
			})
	}

	return m
}
