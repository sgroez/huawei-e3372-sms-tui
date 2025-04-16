package main

import (
	"fmt"
	"time"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	API, err := api.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if smsList, err := API.ReceiveUnreadSms(); err == nil {
				fmt.Println(smsList)
			}
		}
	}()
	select {}
}