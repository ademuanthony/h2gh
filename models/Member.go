package models

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model

	ReferralID int `gorm:"index"`

	Name string
	PhoneNumber string
	Email string
	Password string
	EncryptedPassword string

	BankID uint
	AccountName string
	AccountNumber string

	Bank Bank `gorm:"ForeignKey:BankID;save_associations:false"`

	Payments []Payment
}
