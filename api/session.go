package api

import (
	"bytes"
	"encoding/xml"
	"log"
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

	tokens, err := initialize(session)
	if err != nil {
		return nil, err
	}
	session.verificationTokens = tokens

	return session, nil
}

func initialize(session *session) ([]string, error) {
	tokens := []string{}
	resp, err := get(session, "")
	if err != nil {
		return tokens, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return tokens, err
	}

	doc.Find("meta[name='csrf_token']").Each(func(i int, s *goquery.Selection) {
		token, _ := s.Attr("content")
		tokens = append(tokens, token)
	})
	return tokens, nil
}

func getToken(session *session) string {
	if len(session.verificationTokens) <= 0 {
		log.Println("WARN: No verification token left in list!")
		return ""
	}
	token := session.verificationTokens[0]
	session.verificationTokens = session.verificationTokens[1:]
	return token
}

func get(session *session, endpoint string) (*http.Response, error) {
	url := session.baseUrl + endpoint
	verificationToken := getToken(session)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
    
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

func post(session *session, endpoint string, data any) (*http.Response, error) {
	data_encoded, err := xml.Marshal(data)
	
	if err != nil {
		return nil, err
	}

	url := session.baseUrl + endpoint
	verificationToken := getToken(session)
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