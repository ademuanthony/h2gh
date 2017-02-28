package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/TaxMaster/controllers"
)

func setAccountRoute() {
	beego.Router("/account/create",&controllers.AuthController{},"*:Create")
}
