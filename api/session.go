package api

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/http/cookiejar"
)

type session struct {
	client http.Client
}

func newSession() (*session, error) {
	session := new(session)
	jar, err := cookiejar.New(&cookiejar.Options{})

	if err != nil {
		return nil, err
	}

	client := http.Client{
		Jar: jar,
	}
	session.client = client

	return session, nil
}

func get(session *session, url string, verificationToken string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
    
	req.Header.Set("__RequestVerificationToken", verificationToken)

	resp, err := session.client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func post(session *session, url string, verificationToken string, data any) (*http.Response, error) {
	data_encoded, err := xml.Marshal(data)
	
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

	return resp, nil
}