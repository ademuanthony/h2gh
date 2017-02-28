package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/TaxMaster/controllers"
)

func setSiteRoute() {
	beego.Router("/", &controllers.MainController{})
}
