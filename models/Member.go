package models

import (
	"github.com/willf/pad"
	"strconv"
	"time"
)

type Member struct {
	Id           	int64
	ReferralId	int64
	Email        	string
	Password     	string
	HashPassword 	string
	FirstName    	string
	LastName     	string
	PhoneNumber  	string
	AccountName  	string
	AccountNumber 	string
	Status 		string
	CreatedAt	time.Time

	Bank *Bank `orm:"rel(fk)"`
}

func (this Member) GetReferralCode() string {
	if this.Id == 0{
		return ""
	}
	idString := strconv.FormatInt(this.Id, 10)
	return "H2GH" + pad.Left(idString, 4, "0")
}

func (this Member) GetFromReferralId(referralCode string) (int64, error) {
	idString := referralCode[4:]
	return strconv.ParseInt(idString, 10, 64)

}