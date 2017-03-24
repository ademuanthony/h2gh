package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/TaxMaster/models"
	"net/http"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
)

type PublicPayment struct {
	beego.Controller
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (this *PublicPayment) MakePayment() {
	if this.Ctx.Input.Method() == "POST"{
		var checked bool = true
		flash := beego.NewFlash()
		flash.Store(&this.Controller)

		o := orm.NewOrm()

		taxType := models.TaxType{}

		taxTypeId, _ := this.GetInt("taxTypeId")
		amount, _ := this.GetFloat("amount")
		tin := this.GetString("tin")
		note := this.GetString("note")

		err := o.QueryTable("tax_type").RelatedSel().Filter("id", taxTypeId).One(&taxType)

		if err != nil{
			flash.Error("Invalid Tax Type Selected")
			this.Redirect("/", 302)
		}

		var taxPayment models.PublicPayment
		taxPayment.Amount = amount
		taxPayment.Date = time.Now()
		taxPayment.Note = note
		taxPayment.Tin = tin;
		var linkedTaxType = new(models.TaxType)
		linkedTaxType.Id = taxType.Id
		taxPayment.TaxType = linkedTaxType
		var status = new(models.Status)
		status.Id = models.StatusPending
		taxPayment.Status = status

		_, err = o.Insert(&taxPayment)
		if err != nil{
			flash.Error("Error in processing payment. Please try again later")
			checked = false
		}

		if checked{
			ref := 100000 + taxPayment.Id
			refString := strconv.Itoa(ref)
			this.Ctx.Redirect(302, "/public/payment/pay-online?ref=" + refString)
			return
		}else{
			this.Ctx.Redirect(302, "/")
		}

	}
}

func (this *PublicPayment) PayOnline()  {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	ref, err := this.GetInt("ref")
	if err != nil{
		panic("Invlaid payment reference")
		this.Redirect("/", 302)
	}
	id := ref - 100000

	o := orm.NewOrm()
	taxPayment := models.PublicPayment{}
	err = o.QueryTable("public_payment").RelatedSel().Filter("id", id).One(&taxPayment)

	fmt.Printf("%+v\n", taxPayment)
	if err != nil{
		flash.Error("Error in resolving payment information. Please try again later")
	}

	this.Data["ref"] = ref
	this.Data["Payment"] = taxPayment
	this.TplName = "public/pay-online.html"
}

func (this *PublicPayment) Ping() {
	if this.Ctx.Input.Method() == "POST"{
		ref,_ := this.GetInt("transaction_id")
		id := ref - 100000

		o := orm.NewOrm()
		taxPayment := models.PublicPayment{}
		err := o.QueryTable("public_payment").RelatedSel().Filter("id", id).One(&taxPayment)

		if err != nil{
			panic("Error in resolving payment information from vpay")
		}
		confirmationUrl := "https://voguepay.com/?v_transaction_id=" + string(ref) + "&type=json"

		transaction := models.VoguepayTransaction{}
		getJson(confirmationUrl, transaction)

		if transaction.Status == "Approved"{
			var status = new(models.Status)
			status.Id = models.StatusApproved
			taxPayment.Status = status

			_, _ = o.Update(taxPayment)
		}
	}
}

func (this *PublicPayment) Success()  {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	flash.Success("You have successfully paid your tax. Thank you for performing your obligation")
	this.Redirect("/", 302)
}

func (this *PublicPayment) Fail()  {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	flash.Error("Your payment was not successful. Please try again")
	this.Redirect("/", 302)
}