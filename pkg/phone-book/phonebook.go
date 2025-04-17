package phonebook

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Phonebook struct {
	db *gorm.DB
}

type Conversation struct {
	Phone string `gorm:"primaryKey"`
	Name  string
}

func NewPhonebook() (*Phonebook, error) {
	phonebook := Phonebook{}
	db, err := gorm.Open(sqlite.Open("phonebook.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	phonebook.db = db
	db.AutoMigrate(&Conversation{})

	return &phonebook, nil
}

func (phonebook *Phonebook) FirstOrCreateConversation(phone string) (*Conversation, error) {
	var conversation Conversation
	err := phonebook.db.FirstOrCreate(&conversation, Conversation{Name: phone, Phone: phone}).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (phonebook *Phonebook) FindConversation() ([]Conversation, error) {
	var conversations []Conversation
	err := phonebook.db.Find(&conversations).Error
	if err != nil {
		return nil, err
	}
	return conversations, nil
}