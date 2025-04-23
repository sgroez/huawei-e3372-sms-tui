package huawei_e3372_sms

import (
	"encoding/xml"
	"time"

	"github.com/sgroez/huawei-e3372-sms-tui/internal/helper"
)

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

func (session *Session) SendSms(options SmsSendOptions) error {
	resp, err := session.Post("api/sms/send-sms", options)
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