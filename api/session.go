package api

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
)

type session struct {
	client http.Client
	baseUrl string
	verificationTokens []string
}

func newSession(baseUrl string) (*session, error) {
	session := new(session)
	jar, err := cookiejar.New(&cookiejar.Options{})

	if err != nil {
		return nil, err
	}

	client := http.Client{
		Jar: jar,
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

type Token struct {
	XMLName xml.Name `xml:"request"`
	Token string `xml:"token"`
}

func (session *session) requestToken() (string, error) {
	url := session.baseUrl + "webserver/token"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var token Token
	err = xml.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}

	return token.Token, nil

}

func (session *session) getToken() (string, error) {
	if len(session.verificationTokens) <= 0 {
		return session.requestToken()
	}
	token := session.verificationTokens[0]
	session.verificationTokens = session.verificationTokens[1:]
	return token, nil
}

func (session *session) get(endpoint string) (*http.Response, error) {
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

func (session *session) post(endpoint string, data any) (*http.Response, error) {
	data_encoded, err := xml.Marshal(data)
	
	if err != nil {
		return nil, err
	}

	url := session.baseUrl + endpoint
	verificationToken, err := session.getToken()
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

	//test if this is working
	newToken := resp.Header.Get("__requestverificationtoken")
	if newToken != "" {
		session.verificationTokens = append(session.verificationTokens, newToken)
	} 

	return resp, nil
}