package api

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type Api struct {
	session *session
	baseUrl string
	verificationTokens []string
}

func NewApi(baseUrl string) (*Api, error) {
	api := new(Api)
	session, err := newSession()

	if err != nil {
		return nil, err
	}
	api.session = session
	api.baseUrl = baseUrl

	get(session, baseUrl, "")

	token, err := requestToken(api)

	if err != nil {
		return nil, err
	}

	api.verificationTokens = append(api.verificationTokens, token)

	return api, nil
}

func getToken(api *Api) string {
	if len(api.verificationTokens) <= 0 {
		fmt.Println("WARN: No verification token left in list!")
		return ""
	}
	token := api.verificationTokens[0]
	api.verificationTokens = api.verificationTokens[1:]
	return token
}

func apiGet(api *Api, endpoint string) (*http.Response, error) {
	url := api.baseUrl + endpoint
	token := getToken(api)
	resp, err := get(api.session, url, token)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func apiPost(api *Api, endpoint string, data any) (*http.Response, error) {
	url := api.baseUrl + endpoint
	token := getToken(api)
	resp, err := post(api.session, url, token, data)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func requestToken(api *Api) (string, error) {
	resp, err := apiGet(api, "webserver/token")

	if err != nil {
		return "", err
	}

	var tr TokenResponse 
	err = xml.NewDecoder(resp.Body).Decode(&tr)

	if err != nil {
		return "", err
	}

	return tr.Token, nil
}

func GetSmsList(api *Api) (*SmsListResponse, error) {
	resp, err := apiPost(api, "sms/sms-list", SmsListRequest{
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

	var resp_decoded SmsListResponse
	err = xml.NewDecoder(resp.Body).Decode(&resp_decoded)

	if err != nil {
		return nil, err
	}

	return &resp_decoded, nil
}

func DeleteSms(api *Api, id int) error {
	resp, err := apiPost(api, "sms/delete-sms", SmsDeleteRequest{Index: id})
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

func SendSms(api *Api, numbers []string, message string) error {
	formattedTime := time.Now().Format("2025-04-10 15:25:54")
	phones := []Phone{}
	for _, value := range numbers {
		phones = append(phones, Phone{Phone: value})
	}
	
	resp, err := apiPost(api, "sms/send-sms", SmsSendRequest{
															Index: -1,
															Phones: phones,
															Sca: "",
															Content: message,
															Length: len(message),
															Reserved: 1,
															Date: formattedTime,
	})
	if err != nil {
		return err
	}

	var resp_decoded SimpleResponse
	err = xml.NewDecoder(resp.Body).Decode(&resp_decoded)

	if err != nil {
		return  err
	}
	return nil
}