package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UIConversationList struct {
	*tview.Flex
	primitives []tview.Primitive
	selectedItem int
}

func NewUIConversationList(onFocus func(primitive tview.Primitive)) *UIConversationList {
	conversationList := UIConversationList{primitives: []tview.Primitive{}, selectedItem: -1}
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyDown:
			if conversationList.selectedItem < len(conversationList.primitives) - 1 {
				conversationList.selectedItem++
				onFocus(conversationList.primitives[conversationList.selectedItem])
			}
			return nil
		case tcell.KeyUp:
			if conversationList.selectedItem > 0 {
				conversationList.selectedItem--
				onFocus(conversationList.primitives[conversationList.selectedItem])
			}
			return nil
		}
		return event
	})
	conversationList.Flex = flex
	return &conversationList
}

func (conversationList *UIConversationList) AddConversation(phone string, name string, onEnter func(phone string)) {
		button := tview.NewButton(name).SetSelectedFunc(func() {
			onEnter(phone)
		})
		conversationList.AddItem(button, 1, 0, false)
		conversationList.primitives = append(conversationList.primitives, button)
}