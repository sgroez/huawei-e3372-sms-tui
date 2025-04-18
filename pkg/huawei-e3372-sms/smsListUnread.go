package huaweie3372sms

func(api *Api) SmsListUnread() (*SmsList, error) {
	smsCount, err := api.SmsCount()
	if err != nil {
		return nil, err
	}
	smsListOptions := NewSmsListOptions()
	smsListOptions.ReadCount = smsCount.LocalUnread
	smsListOptions.UnreadPreferred = 1
	smsList, err := api.SmsList(smsListOptions)
	if err != nil {
		return nil, err
	}
	return smsList, nil
}