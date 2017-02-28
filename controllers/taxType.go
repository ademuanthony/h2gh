package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/TaxMaster/models"
	"fmt"
)

type TaxTypeController struct {
	beego.Controller
}

func (this *TaxTypeController) Index() {
	this.TplName = "taxtype/index.html"
	this.Layout = "shared/layout.html";
	this.Data["TaxTypeActive"] = "active"

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
}

func (this *TaxTypeController) Create() {
	if this.Ctx.Input.Method() == "POST"{
		var checked bool
		flash := beego.NewFlash()
		flash.Store(&this.Controller)

		taxType := models.TaxType{}
		name := this.GetString("name")
		taxType.Name = name;

		o := orm.NewOrm()
		_, err := o.Insert(&taxType)
		if err != nil{
			flash.Error("Error in adding type")
			checked = false
		}

		if checked {
			flash.Success("Tax Type Added")
		}

		this.Ctx.Redirect(302, "/taxes/types")
	}
}

func (this *TaxTypeController) Edit() {

}

func (this *TaxTypeController) Delete() {

}
