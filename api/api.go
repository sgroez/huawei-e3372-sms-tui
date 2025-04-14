package api

import (
	"encoding/xml"
	"time"

	"gorm.io/gorm"
)

var TIMESTAMP_LAYOUT = "2006-01-02 15:04:05"

type Api struct {
	db *gorm.DB
	session *session
	User *User
}

func NewApi(baseUrl string, username string, phone string) (*Api, error) {
	api := new(Api)
	session, err := newSession(baseUrl)
	if err != nil {
		return nil, err
	}
	api.session = session

	db, err := newDatabase("chats.db")
	if err != nil {
		return nil, err
	}
	api.db = db

	user, err := firstOrCreateUser(api.db, username, phone)
	if err != nil {
		return nil, err
	}
	api.User = user

	return api, nil
}


func deleteSms(api *Api, id int) error {
	resp, err := post(api.session, "api/sms/delete-sms", SmsDeleteRequest{Index: id})
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

func LoadData(api *Api) {
//implement initial data loading from database
}

func GetMessages(api *Api) ([]Message, error) {
	resp, err := post(api.session, "api/sms/sms-list", SmsListRequest{
															PageIndex: 1,
															ReadCount: 20,
															BoxType: 1,
															SortType: 0,
															Ascending: 0,
															UnreadPreferred: 0,
															})
	if err != nil {
		return nil, err
	}

	var smslist SmsListResponse
	err = xml.NewDecoder(resp.Body).Decode(&smslist)
	if err != nil {
		return nil, err
	}

	messages := []Message{}
	for _, sms := range smslist.Sms {
		conversation, err := prepareConversation(api, sms.Phone)
		if err != nil {
			return nil, err
		}
		timestamp, err := time.Parse(TIMESTAMP_LAYOUT, sms.Date)
		if err != nil {
			return nil, err
		}
		message, err := createMessage(api.db, conversation.ID, conversation.Participant2ID, conversation.Participant1ID, sms.Content, timestamp)
		if err != nil {
			return nil, err
		}
		deleteSms(api, sms.Index)
		messages = append(messages, *message)
	}

	return messages, nil
}

func SendMessage(api *Api, phone string, content string) (*Message, error) {
	timestamp := time.Now()
	formattedTime := timestamp.Format(TIMESTAMP_LAYOUT)
	phones := []Phone{{Phone: phone}}
	
	resp, err := post(api.session, "api/sms/send-sms", SmsSendRequest{
															Index: -1,
															Phones: phones,
															Sca: "",
															Content: content,
															Length: len(content),
															Reserved: 1,
															Date: formattedTime,
	})
	if err != nil {
		return nil, err
	}

	var resp_decoded SimpleResponse
	err = xml.NewDecoder(resp.Body).Decode(&resp_decoded)
	if err != nil {
		return  nil, err
	}
	//check response for error tag in xml

	conversation, err := prepareConversation(api, phone)
	if err != nil {
		return nil, err
	}
	message, err := createMessage(api.db, conversation.ID, conversation.Participant1ID, conversation.Participant2ID, content, timestamp )
	if err != nil {
		return nil, err
	}

	return message, nil
}

func prepareConversation(api *Api, phone string) (*Conversation, error) {
		externalUser, err := firstOrCreateUser(api.db, "", phone)
		if err != nil {
			return nil, err
		}
		conversation, err := firstOrCreateConversation(api.db, api.User.ID, externalUser.ID, "")
		if err != nil {
			return nil, err
		}
		return conversation, nil
}