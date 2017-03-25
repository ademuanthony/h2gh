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

	var maps []orm.Params

	o.Raw("SELECT SUM(amount) as totalPayment FROM public_payment").Values(&maps)
	totalPayment := maps[0]["totalPayment"]

	o.Raw("SELECT SUM(amount) as confirmedPayment FROM public_payment WHERE status_id = ?", models.StatusPending).Values(&maps)
	confirmedPayment := maps[0]["confirmedPayment"]

	o.Raw("SELECT SUM(amount) as pendingPayment FROM public_payment WHERE status_id = ?", models.StatusApproved).Values(&maps)
	pendingPayment := maps[0]["pendingPayment"]


	this.Data["TotalPayment"] = totalPayment
	this.Data["ConfirmedPayment"] = confirmedPayment
	this.Data["PendingPament"] = pendingPayment
	this.TplName = "site/index.html"
}

