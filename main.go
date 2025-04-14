package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	instance, err := api.NewApi("http://192.168.8.1/", os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	args := os.Args[3:]
	if len(args) <= 0 {
		panic(errors.New("Please pass in the method as command line argument!"))
	}
	switch(args[0]) {
	case "list_sms":
		slr, err := api.GetMessages(instance)
		if err != nil {
			panic(err)
		}
		fmt.Println(slr)
	case "send_sms":
		if len(args) <= 2 {
			panic(errors.New("Please pass in at least on phone number and one message in that order when choosing the send_sms method!"))
		}
		phone := args[1]
		message := args[2]
		api.SendMessage(instance, phone, message )
	}
}