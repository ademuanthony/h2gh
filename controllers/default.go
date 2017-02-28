package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Layout = "shared/layout.html"
	this.Data["Website"] = "taxmaster.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["DashboardActive"] = "active"
	this.TplName = "site/index.html"
}
