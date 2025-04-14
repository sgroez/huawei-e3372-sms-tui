package api

import (
	"encoding/xml"
	"time"
)

type SimpleResponse struct {
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
	Sms []SmsResponse `xml:"Messages>Message"`
}

type SmsResponse struct {
	Index   int `xml:"Index"`
    Phone   string `xml:"Phone"`
    Content string `xml:"Content"`
    Date    string `xml:"Date"`
}

type SmsDeleteRequest struct {
	XMLName xml.Name `xml:"request"`
	Index   int   `xml:"Index"`
} 

type SmsSendRequest struct {
	XMLName xml.Name `xml:"request"`
	Index   int   `xml:"Index"`
	Phones  []Phone `xml:"Phones"`
	Sca      string `xml:"Sca"`
	Content  string `xml:"Content"`
	Length   int `xml:"Length"`
	Reserved int `xml:"Reserved"`
	Date     string `xml:"Date"`
}

type Phone struct {
	Phone string `xml:"Phone"`
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Phone string `gorm:"not null;uniqueIndex"`
}

type Conversation struct {
	ID            uint      `gorm:"primaryKey"`
	Participant1ID   uint      `gorm:"not null;"`
    Participant2ID   uint      `gorm:"not null;uniqueIndex"`
	Name          string    `gorm:"not null"`
	LastUpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Participant1 User    `gorm:"foreignKey:Participant1ID"`
	Participant2 User    `gorm:"foreignKey:Participant2ID"`
}

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	ConversationID uint  `gorm:"not null"`
	SenderID   uint      `gorm:"not null"`
	RecipientID uint     `gorm:"not null"`
	Content    string    `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Conversation  Conversation `gorm:"foreignKey:ConversationID"`
	Sender        User         `gorm:"foreignKey:SenderID"`
	Recipient     User         `gorm:"foreignKey:RecipientID"`
}