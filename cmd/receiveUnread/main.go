package main

import (
	"fmt"
	"time"

	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
)

func main() {
	api, err := huaweie3372sms.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if smsList, err := api.SmsListUnread(); err == nil {
				fmt.Println(smsList)
			}
		}
	}()
	select {}
}