package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/ademuanthony/TaxMaster/models"
)

type TaxPayerController struct {
	beego.Controller
}

func (this *TaxPayerController) Index() {
	this.TplName = "taxpayer/index.html"
	this.Layout = "shared/layout.html";
	this.Data["TaxPayerActive"] = "active"

	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	o := orm.NewOrm()
	var taxPayers []*models.TaxPayer
	_, err := o.QueryTable("tax_payer").All(&taxPayers)
	if err != nil{
		fmt.Printf("Error in loading tax payers %s", err)
		flash.Error("Error in loading tax payers. Please try again later")
	}
	this.Data["TaxPayers"] = taxPayers;
}

func (this *TaxPayerController) Create() {
	if this.Ctx.Input.Method() == "POST"{
		var checked bool
		flash := beego.NewFlash()
		flash.Store(&this.Controller)

		taxPayer := models.TaxPayer{}
		name := this.GetString("name")
		tin := this.GetString("tin")
		taxPayer.Name = name;
		taxPayer.Tin = tin

		o := orm.NewOrm()
		_, err := o.Insert(&taxPayer)
		if err != nil{
			flash.Error("Error in adding payer")
			checked = false
		}

		if checked {
			flash.Success("Tax Payer Added")
		}

		this.Ctx.Redirect(302, "/taxes/payers")
	}
}

func (this *TaxPayerController) Edit() {

}
func (this *TaxPayerController) Delete() {

}