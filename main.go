package main

import (
	"github.com/rivo/tview"
	"strconv"
	"tViewExercise/model"
)

func main() {
	app := tview.NewApplication()
	lLay := newListView(app, "menu", true)

	rLay := newBox("rBox Demo", true)

	form := newForm("user info", model.NewUser(), true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(lLay, 30, 20, false).
		AddItem(form, 40, 20, true).
		AddItem(rLay, 60, 80, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

func echo(text string) func() {
	return func() {
		//fmt.Println(text)
	}
}

func stopApp(app *tview.Application) func() {
	return func() {
		app.Stop()
	}
}

func newListView(app *tview.Application, title string, border bool) *tview.List {
	v := tview.NewList().
		AddItem("List item 1", "", 'a', echo("a")).
		AddItem("List item 2", "", 'b', echo("b")).
		AddItem("List item 3", "", 'c', echo("c")).
		AddItem("List item 4", "", 'd', echo("b")).
		AddItem("Quit", "Press to exit", 'q', stopApp(app))
	v.SetBorder(border).SetTitle(title)

	return v
}

func newBox(title string, border bool) *tview.Box {
	v := tview.NewBox().
		SetBorder(border).
		SetTitle(title)

	return v
}

func newForm(title string, u model.User, border bool) *tview.Form {
	t := []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}

	form := tview.NewForm().
		AddDropDown("Title", t, 1, nil).
		AddTextView("Id", strconv.Itoa(u.Id), 20, 1, false, false).
		AddInputField("name", u.Name, 20, nil, nil).
		AddInputField("username", u.Username, 20, nil, nil).
		AddInputField("email", u.Email, 20, nil, nil).
		AddPasswordField("Password", u.Password, 20, '*', nil).
		AddTextArea("note", u.Note, 40, 0, 10, nil).
		AddCheckbox("enable", u.Enable, nil).
		AddButton("Save", nil)
	//AddButton("Quit", func() {
	//	app.Stop()
	//})
	form.SetTitle(title).SetBorder(border)

	return form
}
