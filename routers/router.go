package routers

import (
	/*"github.com/ademuanthony/TaxMaster/controllers"
	"github.com/astaxie/beego"*/
)


func init() {
   // beego.Router("/", &controllers.MainController{})
	setSiteRoute()
	setAuthRoute()

}
