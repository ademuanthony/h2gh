package models

type Member struct {
	Id           	int64
	Email        	string
	Password     	string
	HashPassword 	string
	FirstName    	string
	LastName     	string
	PhoneNumber  	string
	AccountName  	string
	AccountNumber 	string
	Status 		string

	Bank *Bank `orm:"rel(fk)"`
}
