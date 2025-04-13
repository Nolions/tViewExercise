package widget

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewFrame(title string, border bool, content tview.Primitive) *tview.Frame {
	frame := tview.NewFrame(content).
		AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite).
		AddText("Header middle", true, tview.AlignCenter, tcell.ColorWhite).
		AddText("Header right", true, tview.AlignRight, tcell.ColorWhite).
		AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
		AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)
	frame.SetBorder(border)
	frame.SetTitle(title)
	frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	return frame
}
