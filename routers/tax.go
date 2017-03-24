package routers

import (
	"github.com/astaxie/beego"
	"github.com/ademuanthony/TaxMaster/controllers"
)

func setTaxRoutes()  {
	beego.Router("/taxes/types",&controllers.TaxTypeController{},"*:Index")
	beego.Router("/taxes/types/create",&controllers.TaxTypeController{},"*:Create")
	beego.Router("/taxes/types/edit",&controllers.TaxTypeController{},"*:Edit")
	beego.Router("/taxes/types/delete",&controllers.TaxTypeController{},"*:Delete")
}

func setTaxPayerRoutes()  {
	beego.Router("/taxes/payers",&controllers.TaxPayerController{},"*:Index")
	beego.Router("/taxes/payers/create",&controllers.TaxPayerController{},"*:Create")
	beego.Router("/taxes/payers/edit",&controllers.TaxPayerController{},"*:Edit")
	beego.Router("/taxes/payers/delete",&controllers.TaxPayerController{},"*:Delete")
}

func setTaxesRoutes()  {
	beego.Router("/taxes",&controllers.TaxController{},"*:Index")
	beego.Router("/taxes/assessments",&controllers.TaxController{},"*:Index")
	beego.Router("/taxes/assessments/create",&controllers.TaxController{},"*:Create")
	beego.Router("/taxes/assessments/edit",&controllers.TaxController{},"*:Edit")
	beego.Router("/taxes/assessments/delete",&controllers.TaxController{},"*:Delete")
}

func setTaxPaymentRoutes() {
	//beego.Router("/taxes/payment/report",&controllers.TaxPaymentController{},"*:Index")
	beego.Router("/taxes/payment/report",&controllers.TaxPaymentController{},"*:PublicPaymentReport")
	beego.Router("/taxes/payment/pending",&controllers.TaxPaymentController{},"*:Pending")
	beego.Router("/taxes/payment/pay",&controllers.TaxPaymentController{},"*:Pay")
	beego.Router("/taxes/payment/approve",&controllers.TaxPaymentController{},"*:Approve")
	beego.Router("/taxes/payment/reject",&controllers.TaxPaymentController{},"*:Reject")
}