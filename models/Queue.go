package models


type Queue struct {
	Id int64
	Member *Member `orm:"rel(fk)"`
	Amount float64
	SortOrder uint
	Description string
	Status string

}