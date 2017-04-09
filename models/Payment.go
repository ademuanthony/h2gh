package models

import (
	"time"
)

type Payment struct {
	Id int64

	CreatedAt time.Time
	// Id of the user that is receiving the payment
	//MemberId   int64 `orm:"rel(fk)"`
	ToMember     *Member `orm:"rel(fk)"`

	// Id of the user that is making the payment
	//FromMemberId int64 `orm:"rel(fk)"`
	FromMember *Member `orm:"rel(fk)"`

	// Payment amount [5000 for referral bonus and 15000 for rebate]
	Amount       float64

	Description  string

	QueueId      int64 `gorm:"index"`

	Status       string

	PenaltyTime time.Time
}
