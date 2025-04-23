package huawei_e3372_sms

func (session *Session) SmsListGroupedByPhone(unread bool) (map[string][]Sms, error) {
	var smsList *SmsList
	if unread {
		unreadList, err := session.SmsListUnread()
		if err != nil {
			return nil, err
		}
		smsList = unreadList
	}else {
		inOutList, err := session.SmsListInOut()
		if err != nil {
			return nil, err
		}
		smsList = inOutList
	}

	grouped := make(map[string][]Sms)
	for _, sms := range smsList.Sms {
		grouped[sms.Phone] = append(grouped[sms.Phone], sms)
	}
	return grouped, nil
}