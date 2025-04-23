package huawei_e3372_sms

import "encoding/xml"

type SmsSetReadOptions struct {
	XMLName xml.Name `xml:"request"`
	Index   int   `xml:"Index"`
}

func NewSmsSetReadOptions(index int) SmsSetReadOptions {
	return SmsSetReadOptions{
		Index: index,
	}
}

func (session *Session) SmsSetRead(options SmsSetReadOptions) error {
	resp, err := session.Post("api/sms/set-read", options)
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