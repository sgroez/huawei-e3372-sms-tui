package ui

import (
	"log"
	"slices"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sgroez/huawei-e3372-sms-tui/internal/helper"
	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
)


type UIConversation struct {
	*tview.Flex
	width int
	list *tview.Flex
}

func NewUIConversation(phone string, width int, onSend func(phone string, content string) error, onLeave func()) *UIConversation {
	uiConversation := UIConversation{width: width}
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	uiConversation.Flex = flex

	list := tview.NewFlex().SetDirection(tview.FlexRow)
	uiConversation.list = list
	flex.AddItem(list, 0, 1, false)

	input := NewUISmsInput(func(text string) {
		phone := phone 
		date := helper.DateToString(time.Now())
		err := onSend(phone, text)
		if err != nil {
			log.Println(err)
		}
		uiConversation.AddSms([]huaweie3372sms.Sms{{Phone: phone, Content: text, Date: date, Status: 3}})
	})
	flex.AddItem(input, 3, 0, true)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			onLeave()
			return nil
		}
		return event
	})

	return &uiConversation
}

func (conversation *UIConversation) AddSms(sms []huaweie3372sms.Sms) {
	slices.Reverse(sms)
	for _, sms := range sms {
		messageBox := tview.NewTextView().
		SetText(sms.Content).
		SetWrap(true).
		SetTextColor(tcell.ColorWhite)

		spacer := tview.NewBox()

		row := tview.NewFlex().SetDirection(tview.FlexColumn)

		if sms.Status < 3 {
			row.AddItem(messageBox, conversation.width/2, 0, false)
			row.AddItem(spacer, 0, 1, false)
		} else {
			row.AddItem(spacer, 0, 1, false)
			row.AddItem(messageBox, conversation.width/2, 0, false)
		}
		conversation.list.AddItem(row, 0, 1, false)
	}
}