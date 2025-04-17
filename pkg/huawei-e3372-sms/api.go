package huaweie3372sms

import (
	"encoding/xml"

	"github.com/sgroez/huawei-e3372-sms-tui/internal/backend"
)

type Api struct {
	session *backend.Session
}

func NewApi(baseUrl string) (*Api, error) {
	api := new(Api)
	session, err :=backend.NewSession(baseUrl)
	if err != nil {
		return nil, err
	}
	api.session = session

	return api, nil
}

type SimpleResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
}

