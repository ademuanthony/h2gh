package controllers

import "github.com/astaxie/beego"

type DashboardController struct {
	beego.Controller
}

func (this *DashboardController) Index()  {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	this.Data["Title"] = "Dashboard"
	this.Data["DashboardActive"] = "active"

	this.Layout = "shared/layout.html"

	this.TplName = "site/index.html"
}