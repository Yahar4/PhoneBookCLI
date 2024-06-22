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
	form := tview.NewForm()

	form.AddDropDown("Gender", []string{"Male", "Female"}, 0, nil)
	form.AddInputField("Name", "", 20, nil, nil)
	form.AddInputField("Number", "", 14, nil, nil)
	form.AddCheckbox("Age 18+", false, nil)
	form.AddButton("Save", func() {
		name := form.GetFormItemByLabel("Name").(*tview.InputField).GetText()

		phoneNumber := form.GetFormItemByLabel("Number").(*tview.InputField).GetText()

		contact := Contact{
			Name:  name,
			Phone: phoneNumber,
		}

		contactList = append(contactList, &contact)
	})
	form.AddButton("Back", main)

	form.SetBorder(true).SetTitle("Enter Person's Info").SetTitleAlign(tview.AlignCenter)

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
