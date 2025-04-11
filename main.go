package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/sgroez/huawei-e3372-sms-tui/api"
)

func main() {
	instance, err := api.NewApi("http://192.168.8.1/api/")
	if err != nil {
		panic(err)
	}
	args := os.Args[1:]
	if len(args) <= 0 {
		panic(errors.New("Please pass in the method as command line argument!"))
	}
	switch(args[0]) {
	case "list_sms":
		slr, err := api.GetSmsList(instance)
		if err != nil {
			panic(err)
		}
		fmt.Println(slr)
	case "delete_sms":
		if len(args) <= 1 {
			panic(errors.New("Please pass in the id of the message to delete when choosing the delete_sms method!"))
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		api.DeleteSms(instance, id)
	case "send_sms":
		if len(args) <= 2 {
			panic(errors.New("Please pass in at least on phone number and one message in that order when choosing the send_sms method!"))
		}
		phones := args[1:len(args) -1]
		message := args[len(args) -1]
		api.SendSms(instance, phones, message)
	}
}