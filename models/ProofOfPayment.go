package models

import "time"

type ProofOfPayment struct {
	Id int64
	Payment Payment
	Member Member
	CreatedAt time.Time
	UpdatedAt time.Time
	Status string
}
