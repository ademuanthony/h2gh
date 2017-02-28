package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/ademuanthony/TaxMaster/models"
)

type TaxController struct {
	beego.Controller
}

func (this *TaxController) Index() {
	this.TplName = "tax/index.html"
	this.Layout = "shared/layout.html";
	this.Data["TaxAssessmentActive"] = "active"

	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	o := orm.NewOrm()

	//load tax payers for the new assessment
	var taxPayers []*models.TaxPayer
	_, err := o.QueryTable("tax_payer").All(&taxPayers)
	if err != nil{
		fmt.Printf("Error in loading tax payers %s", err)
		flash.Error("Error in loading tax payers. Please try again later")
	}
	this.Data["TaxPayers"] = taxPayers;

	//load tax types for creating new assessment
	var taxTypes []*models.TaxType
	_, err = o.QueryTable("tax_type").All(&taxTypes)
	if err != nil{
		fmt.Printf("Error in loading tax types %s", err)
		flash.Error("Error in loading tax types. Please try again later")
	}
	this.Data["TaxTypes"] = taxTypes;

	//load all assessments
	var taxes []*models.Tax
	_, err = o.QueryTable("tax").RelatedSel().All(&taxes)
	if err != nil{
		fmt.Printf("Error in loading tax assessments %s", err)
		flash.Error("Error in loading tax assessments. Please try again later")
	}
	this.Data["Taxes"] = taxes;

}

func (this *TaxController) Create() {
	if this.Ctx.Input.Method() == "POST"{
		var checked bool
		flash := beego.NewFlash()
		flash.Store(&this.Controller)

		o := orm.NewOrm()

		tax := models.Tax{}

		taxTypeId, _ := this.GetInt("taxTypeId")
		taxPayerId, _ := this.GetInt("taxPayerId")
		var taxType = new(models.TaxType)
		taxType.Id = taxTypeId

		var taxPayer = new(models.TaxPayer)
		taxPayer.Id = taxPayerId

		tax.TaxType = taxType
		tax.TaxPayer = taxPayer

		tax.Amount, _ = this.GetFloat("amount")
		tax.Comment = this.GetString("comment")

		var status = new(models.Status)
		status.Id = models.StatusPending
		tax.Status = status

		_, err := o.Insert(&tax)
		if err != nil{
			flash.Error("Error in adding payer")
			checked = false
		}

		if checked {
			flash.Success("Tax Created")
		}

		this.Ctx.Redirect(302, "/taxes")
	}
}

func (this *TaxController) Edit() {

}

func (this *TaxController) Delete() {

}