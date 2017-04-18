package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/h2gh/controllers"
)

func setSiteRoute() {
	beego.Router("/", &controllers.MainController{})


	beego.Router("/dashboard",&controllers.DashboardController{},"*:Index")

	beego.Router("/payment/confirm",&controllers.PaymentController{},"*:Confirm")


	beego.Router("/downlines",&controllers.Promotion{},"*:Downlines")


	beego.Router("/account",&controllers.AccountController{},"*:Index")

	beego.Router("/account/change-password",&controllers.AccountController{},"*:ChangePassword")

}
