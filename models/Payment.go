package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Payment struct {
	gorm.Model

	// Id of the user that is receiving the payment
	MemberId   uint `gorm:"index"`
	Member     Member

	// Id of the user that is making the payment
	FromMemberId uint

	// Payment amount [5000 for referral bonus and 15000 for rebate]
	Amount       float64

	Description  string

	QueueId      uint `gorm:"index"`

	Status       string

	PenaltyTime time.Time
}
