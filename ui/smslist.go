package ui

import (
	"slices"

	"github.com/rivo/tview"
	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

type UISmsList struct {
	*tview.List
}

func NewUISmsList() *UISmsList {
	list := tview.NewList()
	return &UISmsList{List: list}
}

func (list *UISmsList) AddSms(sms []api.Sms) {
	slices.Reverse(sms)
	for _, sms := range sms {
		list.AddItem(sms.Content, sms.Phone + " " + sms.Date, rune(0), nil)
	}
}