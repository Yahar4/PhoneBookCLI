package main

import (
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

type Contact struct {
	Name  string
	Phone string
}

func quit() {
	app.Stop()
}

func addContact() {

}

func showPeople() {

}

func buildPhoneBook() {

}

func main() {
	list := tview.NewList().
		AddItem("Contacts", "Press to see contacts", 's', buildPhoneBook).
		AddItem("People", "Press to see people", 'p', showPeople).
		AddItem("Add contact", "Press to add someone to list", 'a', addContact).
		AddItem("Quit", "Press to shutdown the app", 'q', quit)

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
