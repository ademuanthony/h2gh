package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/TaxMaster/controllers"
)

func setSiteRoute() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/site", &controllers.MainController{}, "*:Site")


	beego.Router("/public/payment/pay", &controllers.PublicPayment{}, "*:MakePayment")
	beego.Router("/public/payment/pay-online", &controllers.PublicPayment{}, "*:PayOnline")

	beego.Router("/public/payment/ping", &controllers.PublicPayment{}, "*:Ping")
	beego.Router("/public/payment/success", &controllers.PublicPayment{}, "*:Success")
	beego.Router("/public/payment/fail", &controllers.PublicPayment{}, "*:Fail")

}
