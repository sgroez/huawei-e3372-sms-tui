package main

import (
	"time"

	"github.com/rivo/tview"
	"github.com/sgroez/huawei-e3372-sms-tui/api"
	"github.com/sgroez/huawei-e3372-sms-tui/ui"
)

func main() {
	API, err := api.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	title := "SMS CLIENT"

	app := tview.NewApplication()
	uiSmsList := ui.NewUISmsList()

	if smsList, err := API.ReceiveSms(api.NewSmsListOptions()); err == nil {
		uiSmsList.AddSms(smsList.Sms)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if smsList, err := API.ReceiveUnreadSms(); err == nil {
				app.QueueUpdateDraw(func() {
					uiSmsList.AddSms(smsList.Sms)
				})
			}
		}
	}()

	frame := ui.CreateFrame(title, uiSmsList)

	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}