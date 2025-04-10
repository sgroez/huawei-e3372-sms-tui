package api

import (
	"encoding/xml"
)

type TokenResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Token   string   `xml:"token"`
} 

type SmsListRequest struct {
	XMLName xml.Name `xml:"request"`
	PageIndex int `xml:"PageIndex"`
	ReadCount int `xml:"ReadCount"`
	BoxType int `xml:"BoxType"`
	SortType int `xml:"SortType"`
	Ascending int `xml:"Ascending"`
	UnreadPreferred int `xml:"UnreadPreferred"`
}

type SmsListResponse struct {
	XMLName  xml.Name `xml:"response"`
	Text     string   `xml:",chardata"`
	Count    string   `xml:"Count"`
	Messages struct {
		Text    string `xml:",chardata"`
		Message struct {
			Text     string `xml:",chardata"`
			Smstat   string `xml:"Smstat"`
			Index    string `xml:"Index"`
			Phone    string `xml:"Phone"`
			Content  string `xml:"Content"`
			Date     string `xml:"Date"`
			Sca      string `xml:"Sca"`
			SaveType string `xml:"SaveType"`
			Priority string `xml:"Priority"`
			SmsType  string `xml:"SmsType"`
		} `xml:"Message"`
	} `xml:"Messages"`
}

type SmsDeleteRequest struct {
	XMLName xml.Name `xml:"request"`
	Text    string   `xml:",chardata"`
	Index   int   `xml:"Index"`
} 

type SimpleResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
}

type Phone struct {
	Text  string `xml:",chardata"`
	Phone string `xml:"Phone"`
}

type SmsSendRequest struct {
	XMLName xml.Name `xml:"request"`
	Text    string   `xml:",chardata"`
	Index   int   `xml:"Index"`
	Phones  []Phone `xml:"Phones"`
	Sca      string `xml:"Sca"`
	Content  string `xml:"Content"`
	Length   int `xml:"Length"`
	Reserved int `xml:"Reserved"`
	Date     string `xml:"Date"`
}