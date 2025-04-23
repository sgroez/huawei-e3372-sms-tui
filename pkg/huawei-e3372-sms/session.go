package huawei_e3372_sms

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Session struct {
	client http.Client
	baseUrl string
	verificationTokens []string
}

func NewSession(baseUrl string) (*Session, error) {
	session := new(Session)
	jar, err := cookiejar.New(&cookiejar.Options{})

	if err != nil {
		return nil, err
	}

	client := http.Client{
		Jar: jar,
		Timeout: time.Second * 5,
	}
	session.client = client

	session.baseUrl = baseUrl

	session.verificationTokens = []string{}

	resp, err := session.client.Get(session.baseUrl)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("meta[name='csrf_token']").Each(func(i int, s *goquery.Selection) {
		token, _ := s.Attr("content")
		session.verificationTokens = append(session.verificationTokens, token)
	})

	return session, nil
}

type token struct {
	XMLName xml.Name `xml:"request"`
	Token string `xml:"token"`
}

func (session *Session) RequestToken() (string, error) {
	url := session.baseUrl + "webserver/token"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var token token
	err = xml.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Token, nil

}

func (session *Session) GetToken() (string, error) {
	if len(session.verificationTokens) <= 0 {
		return session.RequestToken()
	}
	token := session.verificationTokens[0]
	session.verificationTokens = session.verificationTokens[1:]
	return token, nil
}

type SimpleResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
}

func (session *Session) Get(endpoint string) (*http.Response, error) {
	url := session.baseUrl + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := session.client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (session *Session) Post(endpoint string, data any) (*http.Response, error) {
	data_encoded, err := xml.Marshal(data)
	
	if err != nil {
		return nil, err
	}

	url := session.baseUrl + endpoint
	verificationToken, err := session.GetToken()
	if err != nil {
		return nil, err
	}
	xmlHeader := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	fullXML := append(xmlHeader, data_encoded...)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(fullXML))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("__RequestVerificationToken", verificationToken)

    resp, err := session.client.Do(req)

	if err != nil {
		return nil, err
	}

	newToken := resp.Header.Get("__requestverificationtoken")
	if newToken != "" {
		session.verificationTokens = append(session.verificationTokens, newToken)
	} 

	return resp, nil
}