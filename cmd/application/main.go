package main

import (
	"log"

	"github.com/rivo/tview"
	"github.com/sgroez/huawei-e3372-sms-tui/cmd/application/ui"
	phonebook "github.com/sgroez/huawei-e3372-sms-tui/pkg/phone-book"
)

func main() {
	/*
	api, err := huaweie3372sms.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}
	*/

	phonebook, err := phonebook.NewPhonebook()
	if err != nil {
		panic(err)
	}

	title := "SMS CLIENT"

	app := tview.NewApplication()
	/*
	uiSmsList := ui.NewUISmsList(80)

	if smsList, err := api.SmsListInOut(); err == nil {
		uiSmsList.AddSms(smsList.Sms)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if smsList, err := api.SmsListUnread(); err == nil {
				app.QueueUpdateDraw(func() {
					uiSmsList.AddSms(smsList.Sms)
				})
			}
		}
	}()

	uiSmsInput := ui.NewUISmsInput(func(text string) {
		phone := ""
		date := helper.DateToString(time.Now())
		err := api.SendSms(huaweie3372sms.NewSmsSendOptions(phone, text))
		if err != nil {
			log.Println(err)
		}
		uiSmsList.AddSms([]huaweie3372sms.Sms{{Phone: phone, Content: text, Date: date, Status: 3}})
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(uiSmsList, 0, 1, false). 
	AddItem(uiSmsInput, 3, 0, true)
	*/

	uiConversationList := ui.NewUIConversationList()
	if conversations, err := phonebook.FindConversation(); err == nil {
		uiConversationList.AddConversations(conversations, func(phone string){
			log.Println("open conversation", phone)
		})
	}

	frame := ui.CreateFrame(title, uiConversationList)

	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}