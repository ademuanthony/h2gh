package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/h2gh/controllers"
)

func setAuthRoute() {
	beego.Router("/register",&controllers.AuthController{},"*:Register")
	beego.Router("/login",&controllers.AuthController{},"*:Login")

	beego.Router("/logout",&controllers.AuthController{},"*:Logout")
}
