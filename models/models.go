package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"time"
)


func init() {
	orm.RegisterModel(new(Member), new(Bank))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:0000@/h2gh?charset=utf8", 30)
}

type(
	Status struct {
		Id int
		Text string
	}

	TaxType struct {
		Id int
		Name string
	}

	TaxPayer struct {
		Id int
		Tin string
		Name string
		ContactName string
		ContactEmail string
		ContactPhoneNumber string
	}

	Tax struct {
		Id int
		//TaxTypeId int
		//TaxPayerId int
		Comment string
		Amount float64
		DueDate time.Time
		TaxType *TaxType `orm:"rel(fk)"`
		TaxPayer *TaxPayer `orm:"rel(fk)"`
		Status *Status `orm:"rel(fk)"`
	}

	TaxPayment struct {
		Id int
		//TaxId int
		AmountPaid float64
		Date time.Time
		Tax *Tax `orm:"rel(fk)"`
		Status *Status `orm:"rel(fk)"`
	}

	PublicPayment struct {
		Id int
		Tin string
		Amount float64
		Note string
		Date time.Time
		TaxType *TaxType `orm:"rel(fk)"`
		Status *Status `orm:"rel(fk)"`
	}

	VoguepayTransaction struct {
		Total float64
		Status string
		MerchantId string
	}
)

