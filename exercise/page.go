package exercise

import (
	"fmt"
	"github.com/rivo/tview"
)

func NewPages(
	app *tview.Application,
	labs []string,
	pageCount int,
	fun1 func(*tview.Application),
	fun2 func(pages *tview.Pages, page, pageCount int),
) *tview.Pages {
	pages := tview.NewPages()
	for page := 0; page < pageCount; page++ {
		func(page int) {
			pages.AddPage(fmt.Sprintf("page-%d", page),
				NewModel(app, labs, page, 3, pages, fun1, fun2),
				false,
				page == 0)
		}(page)
	}

	return pages
}
