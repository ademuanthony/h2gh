package data

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/ademuanthony/TaxMaster/utilities"
	"strconv"
)

type PublicPaymentRepository struct {
	Db orm.Ormer
}

func (this PublicPaymentRepository ) GetTotalPaymentAmount() interface{} {

	amountParam := []orm.Params{}
	this.Db.Raw("SELECT SUM(amount) as totalAmount FROM public_payment").Values(&amountParam)

	return amountParam[0]["totalAmount"]
}

func (this PublicPaymentRepository ) GetTotalPaymentAmountByStatus(status int) interface{} {
	amountParam := []orm.Params{}
	this.Db.Raw("SELECT SUM(amount) as totalAmount FROM public_payment WHERE status_id >= ?", status).Values(&amountParam)

	return amountParam[0]["totalAmount"]
}

func (this PublicPaymentRepository ) GetTotalPaymentAmountByTaxType(taxTypeId int) interface{} {
	amountParam := []orm.Params{}
	this.Db.Raw("SELECT SUM(amount) as totalAmount FROM public_payment WHERE tax_type_id >= ?", taxTypeId).Values(&amountParam)

	return amountParam[0]["totalAmount"]
}

func (this PublicPaymentRepository ) GetTotalPaymentAmountByDateRange(startDate time.Time, endDate time.Time) interface{} {
	amountParam := []orm.Params{}
	this.Db.Raw("SELECT SUM(amount) as totalAmount FROM public_payment WHERE date >= ? AND date <= ?", startDate, endDate).Values(&amountParam)

	return amountParam[0]["totalAmount"]
}

func (this PublicPaymentRepository) GetPaymentByMonth(year int, month time.Month) interface{} {
	startTime := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(year, month, utilities.GetLastDayOfTheMonth(year, month), 23, 59, 59, 999, time.UTC)
	return this.GetTotalPaymentAmountByDateRange(startTime, endTime)
}

func (this PublicPaymentRepository) GetPaymentsByMonths(numberOfMonthsPassed int) []map[string]interface{} {
	payments := make([]map[string]interface{}, numberOfMonthsPassed)
	currentDate := time.Now()

	var i int
	sn := 0
	for i = numberOfMonthsPassed -1; i >= 0; i--{
		monthValue := int(currentDate.Month()) - i
		year := currentDate.Year()
		month := time.Month(monthValue)
		if monthValue <= 0{
			monthValue += 12
			month = time.Month(monthValue)
			year -= 1
		}
		payment := make(map[string]interface{})
		payment["amount"] = this.GetPaymentByMonth(year, month)
		payment["color"] = "#" + strconv.Itoa(i%7) + "F45FA"
		payment["month"] = month.String() + " " + strconv.Itoa(year)
		payments[sn] = payment
		sn++
	}

	return payments
}