package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)


type UISmsInput struct {
	*tview.InputField
}

func NewUISmsInput(onEnter func(text string)) *UISmsInput {
	inputField := tview.NewInputField().
	SetLabel("> ").
	SetFieldWidth(0)
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			text := inputField.GetText()
			if text != "" {
				onEnter(text)
				inputField.SetText("")
			}
		}
	})
	return &UISmsInput{InputField: inputField}
}
