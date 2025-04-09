package main

import (
	"fmt"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	lLay := tview.NewList().
		AddItem("List item 1", "", 'a', print("a")).
		AddItem("List item 2", "", 'b', print("b")).
		AddItem("List item 3", "", 'c', print("c")).
		AddItem("List item 4", "", 'd', print("b")).
		AddItem("Quit", "Press to exit", 'q', stopApp(app))

	lLay.SetBorder(true).SetTitle("Databases")

	rLay := tview.NewBox().
		SetBorder(true).
		SetTitle("rBox Demo")

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(lLay, 30, 20, true).
		AddItem(rLay, 90, 80, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func print(text string) func() {
	return func() {
		fmt.Println(text)
	}
}

func stopApp(app *tview.Application) func() {
	return func() {
		app.Stop()
	}
}
