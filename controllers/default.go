package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	this.Data["Title"] = "Home"

	this.Layout = "shared/public_layout.html"
	this.TplName = "public/home.html"
}

func (this *MainController) HowItWorks() {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	this.Data["Title"] = "How it Works"

	this.Layout = "shared/public_layout.html"
	this.TplName = "public/howitworks.html"
}
