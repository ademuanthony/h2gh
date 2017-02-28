package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/ademuanthony/TaxMaster/models"
	"time"
)

type TaxPaymentController struct {
	beego.Controller
}


func (this *TaxPaymentController) Index() {
	this.TplName = "taxpayment/report.html"
	this.Layout = "shared/layout.html";
	this.Data["ReportActive"] = "active"

	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	o := orm.NewOrm()

	//load tax payers for the new assessment
	var taxPayments []*models.TaxPayment
	_, err := o.QueryTable("tax_payment").RelatedSel().All(&taxPayments)
	if err != nil{
		fmt.Printf("Error in loading tax payment reports %s", err)
		flash.Error("Error in loading tax payment reports. Please try again later")
	}
	this.Data["TaxPayments"] = taxPayments;
}

func (this *TaxPaymentController) Pending() {
	this.Data["HideNav"] = true
	this.Layout = "shared/layout.html"
	this.TplName = "taxpayment/pending.html"
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	var tin = this.GetString("tin")
	if tin == ""{
		tin = this.GetSession("tin").(string)
		if tin == ""{
			flash.Error("Please enter your TIN to load your pending Taxes")
			return
		}
	}else{
		this.SetSession("tin", tin)
	}
	var taxPayer models.TaxPayer
	o := orm.NewOrm()
	err := o.QueryTable("tax_payer").Filter("tin", tin).One(&taxPayer)
	if err != nil{
		flash.Error("Invalid TIN")
		return
	}
	var taxes []models.Tax
	/*sql := "SELECT * FROM tax INNER JOIN tax_payer ON tax.tax_payer_id = tax_payer.id " +
		"INNER JOIN tax_type ON tax.tax_type_id = tax_type.id WHERE tax_payer.id = ? AND tax.status_id = ?"
	o.Raw(sql, taxPayer.Id, models.StatusActive).QueryRows(&taxes)*/
	_, err = o.QueryTable("tax").RelatedSel().Filter("tax_payer_id", taxPayer.Id).Filter( "status_id", models.StatusActive).All(&taxes)

	this.Data["TaxPayer"] = taxPayer
	this.Data["Taxes"] = taxes
}

func (this *TaxPaymentController) Pay() {
	if this.Ctx.Input.Method() == "POST"{
		var checked bool = true
		flash := beego.NewFlash()
		flash.Store(&this.Controller)

		o := orm.NewOrm()

		tax := models.Tax{}

		taxId, _ := this.GetInt("taxId")
		amount, _ := this.GetFloat("amount")
		err := o.QueryTable("tax").RelatedSel().Filter("id", taxId).One(&tax)

		if err != nil{
			flash.Error("Invalid Tax Selected")
			this.Redirect("/taxes/payment/pending", 400)
		}
		if amount != tax.Amount{
			flash.Error("Invalid amount")
			this.Redirect("/taxes/payment/pending", 400)
		}

		var taxPayment models.TaxPayment
		taxPayment.AmountPaid = amount
		taxPayment.Date = time.Now()
		var linkedTax = new(models.Tax)
		linkedTax.Id = tax.Id
		taxPayment.Tax = linkedTax
		var status = new(models.Status)
		status.Id = models.StatusPending
		taxPayment.Status = status

		_, err = o.Insert(&taxPayment)
		if err != nil{
			flash.Error("Error in adding payment")
			checked = false
		}

		if checked {
			status.Id = models.StatusPaid
			tax.Status = status
			o.Update(&tax)
			flash.Success("Tax Payment Created")
		}

		this.Ctx.Redirect(302, "/taxes/payment/pending?tin=" + tax.TaxPayer.Tin)
	}
}

func (this *TaxPaymentController) Approve() {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	var taxPayment *models.TaxPayment
	err := o.QueryTable("tax_payment").Filter("id", id).One(&taxPayment)
	if err != nil{
		fmt.Printf("Error in loading tax payment %s", err)
		flash.Error("Error in loading tax payment. Please try again later")
	}
	var status = new(models.Status)
	status.Id = models.StatusApproved
	taxPayment.Status = status
	o.Update(taxPayment)
	flash.Success("Payment Approved")
	this.Redirect("/payment/report", 200)
}

func (this *TaxPaymentController) Reject() {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	var taxPayment *models.TaxPayment
	err := o.QueryTable("tax_payment").Filter("id", id).One(&taxPayment)
	if err != nil{
		fmt.Printf("Error in loading tax payment %s", err)
		flash.Error("Error in loading tax payment. Please try again later")
	}
	var status = new(models.Status)
	status.Id = models.StatusRejected
	taxPayment.Status = status
	o.Update(taxPayment)
	flash.Success("Payment Approved")
	this.Redirect("/payment/report", 200)
}