package models

import "github.com/jinzhu/gorm"

type Queue struct {
	gorm.Model

	MemberId uint `gorm:"index"`
	Amount float64
	SortOrder uint
	Description string
	Status string

	Payment Payment `gorm:"save_associations:false"`
}
