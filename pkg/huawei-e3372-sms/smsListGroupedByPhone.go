package huaweie3372sms

func (api *Api) SmsListGroupedByPhone(unread bool) (map[string][]Sms, error) {
	var smsList *SmsList
	if unread {
		unreadList, err := api.SmsListUnread()
		if err != nil {
			return nil, err
		}
		smsList = unreadList
	}else {
		inOutList, err := api.SmsListInOut()
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