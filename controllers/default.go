package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/TaxMaster/models"
	"fmt"
	"encoding/json"
	"html/template"
	"strings"
	"github.com/ademuanthony/TaxMaster/data"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	flash := beego.NewFlash()
	flash.Store(&this.Controller)
	o := orm.NewOrm()
	var taxTypes []*models.TaxType
	_, err := o.QueryTable("tax_type").All(&taxTypes)
	if err != nil{
		fmt.Printf("Error in loading tax types %s", err)
		flash.Error("Error in loading tax types. Please try again later")
	}
	this.Data["TaxTypes"] = taxTypes;

	this.TplName = "public/index.html"
}

func (this *MainController) Site() {
	this.Layout = "shared/layout.html"

	o := orm.NewOrm()
	publicPaymentRepository := data.PublicPaymentRepository{Db: o}

	totalPayment := publicPaymentRepository.GetTotalPaymentAmount()

	pendingPayment := publicPaymentRepository.GetTotalPaymentAmountByStatus(models.StatusPending)
	fmt.Printf("pendingPayment: %+v\n", pendingPayment)

	confirmedPayment := publicPaymentRepository.GetTotalPaymentAmountByStatus(models.StatusApproved)


	var taxTypes []*models.TaxType
	_, err := o.QueryTable("tax_type").All(&taxTypes)
	if err != nil{
		fmt.Printf("Error in loading tax types %s", err)
	}

	items := make([]map[string]interface{}, len(taxTypes))
	title := make(map[string]interface{})
	title["Type"] = "Amount"
	items[0] = title

	jsonString := "[['Tax_Type', 'Amount']";

	for _, taxType := range taxTypes{
		item := make(map[string]interface{})
		amount := publicPaymentRepository.GetTotalPaymentAmountByTaxType(taxType.Id)
		item[taxType.Name] = amount

		jsonString += "," + buildJsArrayString(item)

	}
	jsonString += "]"


	//fmt.Printf("jString: %+v\n", jsonString)
	//get the payment by months for graph
	paymentsByMonth := publicPaymentRepository.GetPaymentsByMonths(12)
	//fmt.Printf("%v\n", paymentsByMonth)
	paymentsByMonthJson, _ := json.Marshal(paymentsByMonth)

	this.Data["TotalPayment"] = totalPayment
	this.Data["ConfirmedPayment"] = confirmedPayment
	this.Data["PendingPayment"] = pendingPayment
	this.Data["ChartData"] = template.JS(jsonString)
	this.Data["MothlyChartData"] = template.JS(paymentsByMonthJson)
	this.TplName = "site/index.html"
}

func buildJsArrayString(items interface{}) string {
	jsonData, err := json.Marshal(items)
	if err != nil{
		return ""
	}
	stringPaths := strings.Split(string(jsonData), ":")
	first := stringPaths[0]
	second := stringPaths[1]

	second = strings.Replace(second, "\"", "", -1)

	jString := first + "," + second

	jString = strings.Replace(jString, "{", "[", -1)
	jString = strings.Replace(jString, "}", "]", -1)
	jString = strings.Replace(jString, ":", ",", -1)
	jString = strings.Replace(jString, "null", "0", -1)
	return jString
}

func buildJsArrayStringFromMap(items map[string]interface{}) string {
	jsonSting := "[['Month', 'Amount']"
	for key, value := range items{
		itemString := "['" + string(key) + "', " + value.(string)+ "]"
		jsonSting += ", " + itemString
	}
	return jsonSting + "]"
}