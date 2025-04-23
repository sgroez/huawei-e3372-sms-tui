package main

import (
	"fmt"

	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
)

func main() {
	api, err := huaweie3372sms.NewSession("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	smsListOptions := huaweie3372sms.NewSmsListOptions()

	smslist, err := api.SmsList(smsListOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println(smslist)
}