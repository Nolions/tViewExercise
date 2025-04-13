package exercise

import (
	"github.com/rivo/tview"
	"tViewExercise/model"
)

func NewForm(title string, u model.User, border bool) *tview.Form {
	t := []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}

	form := tview.NewForm().
		AddDropDown("Title", t, 1, nil).
		//AddTextView("Id", strconv.Itoa(u.Id), 20, 1, false, false).
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
