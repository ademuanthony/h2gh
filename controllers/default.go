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
	this.Data["Website"] = "taxmaster.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["DashboardActive"] = "active"
	this.TplName = "site/index.html"
}

