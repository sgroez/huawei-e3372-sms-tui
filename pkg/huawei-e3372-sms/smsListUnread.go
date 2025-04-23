package huawei_e3372_sms

func(session *Session) SmsListUnread() (*SmsList, error) {
	smsCount, err := session.SmsCount()
	if err != nil {
		return nil, err
	}
	smsListOptions := NewSmsListOptions()
	smsListOptions.ReadCount = smsCount.LocalUnread
	smsListOptions.UnreadPreferred = 1
	smsList, err := session.SmsList(smsListOptions)
	if err != nil {
		return nil, err
	}
	return smsList, nil
}