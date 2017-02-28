package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/TaxMaster/controllers"
)

func setAuthRoute() {
	beego.Router("/auth/login",&controllers.AuthController{},"*:Login")

	beego.Router("/auth/logout",&controllers.AuthController{},"*:Logout")
}
