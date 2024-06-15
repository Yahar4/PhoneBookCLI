package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

type Contact struct {
	Name  string
	Phone string
}

var contactList = []*Contact{
	{
		Name:  "John",
		Phone: "+11233216767",
	},
}

func quit() {
	app.Stop()
}

func showPeople() {
	list := tview.NewList()
	list.SetBorder(true).SetTitle("Contacts (press q to quit)")

	for _, contact := range contactList {
		list.AddItem(contact.Name, contact.Phone, 'q', main)
	}

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func addContact() {
	form := tview.NewForm().
		AddDropDown("Gender", []string{"Male", "Female"}, 0, nil).
		AddInputField("Name", "", 20, nil, nil).
		AddCheckbox("Age 18+", false, nil).
		AddButton("Save", nil).
		AddButton("Back", main)

	form.SetBorder(true).SetTitle("Enter Person's Info").SetTitleAlign(tview.AlignCenter)

	name := form.GetFormItem(2).(*tview.InputField).GetText()

	contact := Contact{Name: name}

	contactList = append(contactList, &contact)

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func main() {
	list := tview.NewList().
		AddItem("People", "Press to see people", 'p', showPeople).
		AddItem("Add contact", "Press to add someone to list", 'a', addContact).
		AddItem("Quit", "Press to shutdown the app", 'q', quit)
	list.SetBorder(true).SetTitle("PhoneBook")

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
