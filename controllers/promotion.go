package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"net/http"
)

type Promotion struct {
	beego.Controller
}

func (this *Promotion) Downlines() {
	this.Data["Title"] = "My Referrals"

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

	var downlines []models.Member
	o.QueryTable(new(models.Member)).Filter("referral_id", currentUserId).All(&downlines)

	this.Data["Downlines"] = downlines

	this.Data["DownlinesActive"] = "active"

	this.Layout = "shared/layout.html"

	this.TplName = "promotion/downlines.html"
}
