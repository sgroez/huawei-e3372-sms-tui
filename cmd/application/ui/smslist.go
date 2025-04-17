package ui

import (
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
)

type UISmsList struct {
	*tview.Flex
	width int
}

func NewUISmsList(width int) *UISmsList {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	return &UISmsList{Flex: flex, width: width}
}

func (smsList *UISmsList) AddSms(sms []huaweie3372sms.Sms) {
	slices.Reverse(sms)
	for _, sms := range sms {
		messageBox := tview.NewTextView().
		SetText(sms.Content).
		SetWrap(true).
		SetTextColor(tcell.ColorWhite)

		spacer := tview.NewBox()

		row := tview.NewFlex().SetDirection(tview.FlexColumn)

		if sms.Status < 3 {
			row.AddItem(messageBox, smsList.width/2, 0, false)
			row.AddItem(spacer, 0, 1, false)
		} else {
			row.AddItem(spacer, 0, 1, false)
			row.AddItem(messageBox, smsList.width/2, 0, false)
		}
		smsList.AddItem(row, 0, 1, false)
	}
}