package api

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newDatabase(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
    if err != nil {
		return nil, err
    }

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Conversation{})
	db.AutoMigrate(&Message{})

	return db, nil
}

func firstOrCreateUser(db *gorm.DB, name string, phone string) (*User, error) {
	user := User{Name: name, Phone: phone}
	err := db.FirstOrCreate(&user, user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func firstOrCreateConversation(db *gorm.DB, localUserID uint, externalUserID uint, name string) (*Conversation, error) {
	conversation := Conversation{Participant1ID: localUserID, Participant2ID: externalUserID, Name: name}
	err := db.FirstOrCreate(&conversation, conversation).Error
	if err != nil {
		return nil, err
	}
	return &conversation, nil 
}

func createMessage(db *gorm.DB, conversationID uint, senderID uint, recipientID uint, content string, timestamp time.Time) (*Message, error) {
	message := Message{ConversationID: conversationID, SenderID: senderID, RecipientID: recipientID, Content: content, Timestamp: timestamp}
	err := db.FirstOrCreate(&message, message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}