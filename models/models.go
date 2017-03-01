package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
	"time"
)

var(
	StatusActive int = 1
	StatusInactive int = 2
	StatusPending int = 3
	StatusApproved int = 4
	StatusRejected int = 5
	StatusPaid = 6
)

func init() {
	orm.RegisterModel(new(UserAuth), new(TaxType), new(TaxPayer), new(Tax), new(TaxPayment), new(Status))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:0000@/taxmaster?charset=utf8", 30)
}

type(
	UserAuth struct {
		Id int
		Email string
		Password string
		HashPassword string
		FirstName string
		LastName string
	}

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
)

