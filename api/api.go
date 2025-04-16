package api

import (
	"encoding/xml"
	"time"

	"github.com/sgroez/huawei-e3372-sms-tui/helper"
)

type Api struct {
	session *session
}

func NewApi(baseUrl string) (*Api, error) {
	api := new(Api)
	session, err := newSession(baseUrl)
	if err != nil {
		return nil, err
	}
	api.session = session

	return api, nil
}

type SimpleResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
}

type SmsCount struct {
    XMLName      xml.Name `xml:"response"`
    LocalUnread  int      `xml:"LocalUnread"`
    LocalInbox   int      `xml:"LocalInbox"`
    LocalOutbox  int      `xml:"LocalOutbox"`
}

func (api *Api) SmsCount() (*SmsCount, error) {
	resp, err := api.session.get("api/sms/sms-count")
	if err != nil {
		return nil, err
	}

	var smsCount SmsCount
	err = xml.NewDecoder(resp.Body).Decode(&smsCount)
	if err != nil {
		return nil, err
	}
	return &smsCount, nil
}

type SmsSetReadOptions struct {
	XMLName xml.Name `xml:"request"`
	Index   int   `xml:"Index"`
}

func NewSmsSetReadOptions(index int) SmsSetReadOptions {
	return SmsSetReadOptions{
		Index: index,
	}
}


func (api *Api) SmsSetRead(options SmsSetReadOptions) error {
	resp, err := api.session.post("api/sms/set-read", options)
	if err != nil {
		return err
	}

	var resp_decoded SimpleResponse
	err = xml.NewDecoder(resp.Body).Decode(&resp_decoded)
	if err != nil {
		return err
	}

	return nil
}

type SmsListOptions struct {
	XMLName xml.Name `xml:"request"`
	PageIndex int `xml:"PageIndex"`
	ReadCount int `xml:"ReadCount"`
	BoxType int `xml:"BoxType"`
	SortType int `xml:"SortType"`
	Ascending int `xml:"Ascending"`
	UnreadPreferred int `xml:"UnreadPreferred"`
}

func NewSmsListOptions() SmsListOptions {
	return SmsListOptions{
		PageIndex: 1,
		ReadCount: 20,
		BoxType: 1,
		SortType: 0,
		Ascending: 0,
		UnreadPreferred: 0,
	}
}

type SmsList struct {
	Sms []Sms `xml:"Messages>Message"`
}

type Sms struct {
	Index   int `xml:"Index"`
    Phone   string `xml:"Phone"`
    Content string `xml:"Content"`
    Date    string `xml:"Date"`
}

func (api *Api) ReceiveSms(options SmsListOptions) (*SmsList, error) {
	resp, err := api.session.post("api/sms/sms-list", options)
	if err != nil {
		return nil, err
	}

	var smslist SmsList
	err = xml.NewDecoder(resp.Body).Decode(&smslist)
	if err != nil {
		return nil, err
	}

	for _, sms := range smslist.Sms {
		api.SmsSetRead(NewSmsSetReadOptions(sms.Index))
	}
	return &smslist, nil
}

func(api *Api) ReceiveUnreadSms() (*SmsList, error) {
	smsCount, err := api.SmsCount()
	if err != nil {
		return nil, err
	}
	smsListOptions := NewSmsListOptions()
	smsListOptions.ReadCount = smsCount.LocalUnread
	smsListOptions.UnreadPreferred = 1
	smsList, err := api.ReceiveSms(smsListOptions)
	if err != nil {
		return nil, err
	}
	return smsList, nil
}

type SmsSendOptions struct {
	XMLName xml.Name `xml:"request"`
	Index   int   `xml:"Index"`
	Phones  []Phone `xml:"Phones"`
	Sca      string `xml:"Sca"`
	Content  string `xml:"Content"`
	Length   int `xml:"Length"`
	Reserved int `xml:"Reserved"`
	Date     string `xml:"Date"`
}

type Phone struct {
	Phone string `xml:"Phone"`
}

func NewSmsSendOptions(phone string, content string) SmsSendOptions {
	return SmsSendOptions{
		Index: -1,
		Phones: []Phone{{Phone: phone}},
		Sca: "",
		Content: content,
		Length: len(content),
		Reserved: 1,
		Date: helper.DateToString(time.Now()),
	}
}

func (api *Api) SendSms(options SmsSendOptions) error {
	resp, err := api.session.post("api/sms/send-sms", options)
	if err != nil {
		return err
	}

	var resp_decoded SimpleResponse
	err = xml.NewDecoder(resp.Body).Decode(&resp_decoded)
	if err != nil {
		return  err
	}
	//check response for error tag in xml
	return nil
}