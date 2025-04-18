package main

import (
	"time"

	"github.com/rivo/tview"
	"github.com/sgroez/huawei-e3372-sms-tui/cmd/application/ui"
	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
	phonebook "github.com/sgroez/huawei-e3372-sms-tui/pkg/phone-book"
)

func main() {
	api, err := huaweie3372sms.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	phonebook, err := phonebook.NewPhonebook()
	if err != nil {
		panic(err)
	}

	title := "SMS CLIENT"
	app := tview.NewApplication()
	pages := tview.NewPages()
	conversationMap := map[string]*ui.UIConversation{}

	smsListGrouped, err := api.SmsListGroupedByPhone(false)
	if err != nil {
		panic(err)
	}

	uiConversationList := ui.NewUIConversationList(func(primitive tview.Primitive) {
		app.SetFocus(primitive)
	})
	for _, group := range smsListGrouped {
		phone := group[0].Phone
		name := phone
		if contact, err := phonebook.FindWithPhone(phone); err == nil {
			name = contact.Name
		}
		uiConversationList.AddConversation(phone, name, func(phone string) {
			pages.SwitchToPage(phone)
		})
	}
	
	pages.AddPage("conversations", uiConversationList, true, true)

	//add function to add contact

	for _, group := range smsListGrouped {
		phone := group[0].Phone
		uiConversation := ui.NewUIConversation(phone, 80, func(phone string, content string) error {
			return api.SendSms(huaweie3372sms.NewSmsSendOptions(phone, content))
		},func() {
			pages.SwitchToPage("conversations")
		})
		uiConversation.AddSms(group)
		pages.AddPage(phone, uiConversation, true, false)
		conversationMap[phone] = uiConversation
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			smsListGrouped, err := api.SmsListGroupedByPhone(true) 
			if err != nil {
				panic(err)
			}
			app.QueueUpdateDraw(func() {
				for _, group := range smsListGrouped {
					phone := group[0].Phone
					name := phone
					if contact, err := phonebook.FindWithPhone(phone); err == nil {
						name = contact.Name
					}
					if pages.HasPage(phone) {
						uiConversation := conversationMap[phone]
						uiConversation.AddSms(group)
					}else {
						uiConversationList.AddConversation(phone, name, func(phone string) {
							pages.SwitchToPage(phone)
						})
						uiConversation := ui.NewUIConversation(phone, 80, func(phone string, content string) error {
							return api.SendSms(huaweie3372sms.NewSmsSendOptions(phone, content))
						},func() {
							pages.SwitchToPage("conversations")
						})
						uiConversation.AddSms(group)
						pages.AddPage(phone, uiConversation, true, false)
						conversationMap[phone] = uiConversation
					}

				}
			})
		}
	}()

	frame := ui.CreateFrame(title, pages)

	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}