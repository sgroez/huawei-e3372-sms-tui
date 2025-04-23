package huawei_e3372_sms

import (
	"encoding/xml"
)

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
	Status int `xml:"Smstat"`
}

func (session *Session) SmsList(options SmsListOptions) (*SmsList, error) {
	resp, err := session.Post("api/sms/sms-list", options)
	if err != nil {
		return nil, err
	}

	var smslist SmsList
	err = xml.NewDecoder(resp.Body).Decode(&smslist)
	if err != nil {
		return nil, err
	}

	for _, sms := range smslist.Sms {
		session.SmsSetRead(NewSmsSetReadOptions(sms.Index))
	}
	return &smslist, nil
}