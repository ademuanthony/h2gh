package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/TaxMaster/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	o := orm.NewOrm()
	var taxTypes []*models.TaxType
	_, err := o.QueryTable("tax_type").All(&taxTypes)
	if err != nil{
		fmt.Printf("Error in loading tax types %s", err)
		flash.Error("Error in loading tax types. Please try again later")
	}
	this.Data["TaxTypes"] = taxTypes;

	this.TplName = "public/index.html"
}

func (this *MainController) Site() {
	this.Layout = "shared/layout.html"
	o := orm.NewOrm()

	var totalPayment float64
	o.Raw("SELECT SUM(amount) as totalPayment FROM public_payment").QueryRow(&totalPayment)

	var confirmedPayment float64
	o.Raw("SELECT COUNT(*) as confirmedPayment FROM public_payment WHERE status_id = 3").QueryRow(&confirmedPayment)

	var pendingPayment float64
	o.Raw("SELECT SUM(amount) as pendingPayment FROM public_payment WHERE status_id = 4").QueryRow(&pendingPayment)


	this.Data["TotalPayment"] = totalPayment
	this.Data["ConfirmedPayment"] = confirmedPayment
	this.Data["PendingPament"] = pendingPayment
	this.TplName = "site/index.html"
}

