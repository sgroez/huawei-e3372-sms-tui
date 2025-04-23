package main

import (
	"errors"
	"fmt"
	"os"

	huaweie3372sms "github.com/sgroez/huawei-e3372-sms-tui/pkg/huawei-e3372-sms"
)

func main() {
	if len(os.Args) < 3 {
		panic(errors.New("Please pass in the following arguments <receiver_phone> <message>!"))
	}
	api, err := huaweie3372sms.NewSession("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	err = api.SendSms(huaweie3372sms.NewSmsSendOptions(os.Args[1], os.Args[2]))
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully send message.")
}