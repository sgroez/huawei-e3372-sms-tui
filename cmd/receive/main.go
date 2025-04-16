package main

import (
	"fmt"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	API, err := api.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	smsListOptions := api.NewSmsListOptions()
	smsListOptions.ReadCount = 1
	smsListOptions.UnreadPreferred = 1

	smslist, err := API.ReceiveSms(smsListOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println(smslist)
}