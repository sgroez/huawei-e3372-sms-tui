package ui

import (
	"github.com/rivo/tview"
	phonebook "github.com/sgroez/huawei-e3372-sms-tui/pkg/phone-book"
)

type UIConversationList struct {
	*tview.Flex
	width int
}

func NewUIConversationList() *UIConversationList {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	return &UIConversationList{Flex: flex}
}

func (conversationList *UIConversationList) AddConversations(conversations []phonebook.Conversation, onEnter func(phone string)) {
	for _, conversation := range conversations {
		button := tview.NewButton(conversation.Name).SetSelectedFunc(func() {
			onEnter(conversation.Phone)
		})
		conversationList.AddItem(button, 1, 0, false)
	}
}