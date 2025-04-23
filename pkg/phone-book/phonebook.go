package phonebook

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Phonebook struct {
	db *gorm.DB
}

type Contact struct {
	Phone string `gorm:"primaryKey"`
	Name  string
}

func NewPhonebook() (*Phonebook, error) {
	phonebook := Phonebook{}
	db, err := gorm.Open(sqlite.Open("phonebook.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	phonebook.db = db
	db.AutoMigrate(&Contact{})

	return &phonebook, nil
}

func (phonebook *Phonebook) FirstOrCreateContact(phone string) (*Contact, error) {
	var contact Contact
	err := phonebook.db.FirstOrCreate(&contact, Contact{Name: phone, Phone: phone}).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (phonebook *Phonebook) UpdateContactName(phone string, name string) {
	phonebook.db.Model(&Contact{}).
    Where("phone = ?", phone).
    Update("name", name)
}

func (phonebook *Phonebook) FindWithPhone(phone string) (*Contact, error) {
	var contact Contact
	err := phonebook.db.First(&contact, phone).Error
	if err != nil {
		return nil, err
	}
	return &contact, nil
}
