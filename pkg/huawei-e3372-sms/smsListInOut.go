package huawei_e3372_sms

import (
	"log"

	"github.com/sgroez/huawei-e3372-sms-tui/internal/helper"
)

func(session *Session) SmsListInOut() (*SmsList, error) {
	smsListOptions := NewSmsListOptions()
	smsListIn, err := session.SmsList(smsListOptions)
	if err != nil {
		return nil, err
	}

	smsListOptions.BoxType = 2
	smsListOut, err := session.SmsList(smsListOptions)
	if err != nil {
		return nil, err
	}

	
	smsListCombined := SmsList{Sms: mergeSortedSmsByDate(smsListIn.Sms, smsListOut.Sms)}
	return &smsListCombined, nil
}

func mergeSortedSmsByDate(a, b []Sms) []Sms {
	result := make([]Sms, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		dateA, errA := helper.StringToDate(a[i].Date)
		dateB, errB := helper.StringToDate(b[j].Date)

		if errA != nil{
			log.Println(errA)
			continue
		}

		if errB != nil {
			log.Println(errB)
			continue
		}

		if dateA.After(dateB) || dateA.Equal(dateB) {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}