package huaweie3372sms

import "encoding/xml"


type SmsCount struct {
    XMLName      xml.Name `xml:"response"`
    LocalUnread  int      `xml:"LocalUnread"`
    LocalInbox   int      `xml:"LocalInbox"`
    LocalOutbox  int      `xml:"LocalOutbox"`
}

func (api *Api) SmsCount() (*SmsCount, error) {
	resp, err := api.session.Get("api/sms/sms-count")
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