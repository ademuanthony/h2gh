package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"github.com/ademuanthony/h2gh/models"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/services"
)

type DashboardController struct {
	beego.Controller
}

func (this *DashboardController) Index()  {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	id := this.Ctx.Input.Session("uid")
	currentUserId, ok := id.(int64)
	if !ok{
		flash.Error("Please login to continue")
		this.Redirect("/login", http.StatusForbidden)
		return
	}

	o := orm.NewOrm()
	accountService := services.AccountService{O:o}

	currentMember, _ := accountService.GetMemberById(currentUserId)

	this.Data["CurrentMember"] = currentMember;
	this.Data["ReferralCode"] = currentMember.GetReferralCode()

	var pendingPayments []models.Payment
	_,err := o.QueryTable(new(models.Payment)).Filter("from_member_id", currentUserId).Filter("status", models.StatusPending).RelatedSel().All(&pendingPayments)

	if err != nil{
		panic(err)
	}
	this.Data["PendingPayments"] = pendingPayments

	var pendingConfirmations []models.Payment
	o.QueryTable(new(models.Payment)).Filter("to_member_id", currentUserId).Filter("status", models.StatusPending).RelatedSel().All(&pendingConfirmations)

	this.Data["PendingConfirmations"] = pendingConfirmations

	var paymentsReceived []models.Payment
	o.QueryTable(new(models.Payment)).Filter("to_member_id", currentUserId).Filter("status", models.StatusConfirmed).
		RelatedSel("FromMember").OrderBy("-id").All(&paymentsReceived)

	this.Data["PaymentsReceived"] = paymentsReceived


	sql := "SELECT SUM(amount) AS totalAmount from payment WHERE from_member_id = ? AND status = ?"
	var params []orm.Params

	o.Raw(sql, currentUserId, models.StatusConfirmed).Values(&params)
	this.Data["TotalDonation"] = params[0]["totalAmount"]


	sql = "SELECT SUM(amount) AS totalAmount from payment WHERE to_member_id = ? AND status = ? AND amount = ?"

	o.Raw(sql, currentUserId, models.StatusConfirmed, 4000).Values(&params)
	this.Data["TotalRebate"] = params[0]["totalAmount"]

	o.Raw(sql, currentUserId, models.StatusConfirmed, 1000).Values(&params)
	this.Data["TotalBonus"] = params[0]["totalAmount"]


	this.Data["Title"] = "Dashboard"
	this.Data["DashboardActive"] = "active"

	this.Layout = "shared/layout.html"

	this.TplName = "dashboard/index.html"
}
