package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	if len(os.Args) < 3 {
		panic(errors.New("Please pass in the following arguments <receiver_phone> <message>!"))
	}
	API, err := api.NewApi("http://192.168.8.1/")
	if err != nil {
		panic(err)
	}

	err = API.SendSms(api.NewSmsSendOptions(os.Args[1], os.Args[2]))
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully send message.")
}