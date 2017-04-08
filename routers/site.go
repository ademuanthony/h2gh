package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/h2gh/controllers"
)

func setSiteRoute() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/dashboard",&controllers.DashboardController{},"*:Index")

}
