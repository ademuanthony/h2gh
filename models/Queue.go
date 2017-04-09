package models

import "time"

type Queue struct {
	Id          int64
	Member      *Member `orm:"rel(fk)"`
	Amount      float64
	SortOrder   int64
	Description string
	Status      string

	CreatedAt   time.Time
	UpdatedAt   time.Time
}