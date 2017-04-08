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

	this.TplName = "public/home.html"
}
