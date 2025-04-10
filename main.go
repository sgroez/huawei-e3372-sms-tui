package main

import (
	"fmt"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	instance, err := api.NewApi("http://192.168.8.1/api/")
	if err != nil {
		panic(err)
	}
	slr, err := api.GetSmsList(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(slr)
	//api.DeleteSms(instance, id)
	//api.SendSms(instance, []string{"number"}, "test message from golang land")
}