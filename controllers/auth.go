package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"strconv"
	"github.com/ademuanthony/h2gh/services"
	"github.com/ademuanthony/h2gh/dto"
	"time"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Login()  {
	this.TplName = "auth/login.html";
	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		email := this.GetString("email")
		password := this.GetString("password");
		o := orm.NewOrm()
		user := models.Member{}

		err := o.QueryTable("member").Filter("email", email).Filter("password", password).One(&user)
		if err == orm.ErrMultiRows {
			flash.Error("Invalid credetial " + email)
			flash.Store(&this.Controller)
			return
		}
		if err == orm.ErrNoRows {
			flash.Error("Invalid credetial " + email)
			flash.Store(&this.Controller)
			return
		}
		
		this.SetSession("uid", user.Id)
			this.Redirect("/dashboard", 302)
			return


		if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password)); err == nil{
			this.SetSession("uid", user.Id)
			this.Redirect("/dashboard", 302)
			return
		}

		flash.Error("Invalid credetial p")
		flash.Store(&this.Controller)
	}
}

func (this *AuthController) Logout() {
	this.DelSession("uid")
	this.Redirect("/login", 308)//removed
}


func (this *AuthController) Register() {
	o := orm.NewOrm()
	o.Begin()

	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	var banks []models.Bank
	_, err := o.QueryTable(new(models.Bank)).All(&banks)
	if err != nil{
		fmt.Printf("%v\n", err)
	}

	this.Data["Banks"] = buildBankOptions(0, banks)
	this.Data["Title"] = "Create Account"

	//fmt.Printf("Banks: %v\n", this.Data["Banks"])

	//get referral code url
	this.Data["ReferralCode"] = this.GetString("ReferralCode")

	this.TplName = "auth/register.html";

	if this.Ctx.Input.Method() == "POST"{
		user := models.Member{}
		if err := this.ParseForm(&user); err != nil {
			//handle error
			fmt.Printf("NUM: ERR: %v\n", err)
		}

		bankId, err := this.GetInt64("BankId")
		if err != nil{
			fmt.Printf("NUM: ERR: %v\n", err)
			flash.Error("Please select your bank")
			this.Data["Member"] = user
			return
		}

		this.Data["Banks"] = buildBankOptions(bankId, banks)


		if this.GetString("ConfirmPassword") != user.Password{
			fmt.Printf("NUM: ERR: %v\n", err)
			flash.Error("Password and Confirm password must be the same")
			this.Data["Member"] = user
			return
		}

		//check email duplicate
		if o.QueryTable(new(models.Member)).Filter("email", user.Email).Exist() {
			flash.Error("The selected email has already been taken");
			this.Data["Member"] = user;
			return
		}

		/*hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("NUM: ERR: %v\n", err)
		} */

		referralCode := this.GetString("ReferralCode")
		this.Data["ReferralCode"] = referralCode

		if referralCode != ""{
			if len(referralCode) != 8{
				flash.Error("Invalid referral Code")
				this.Data["Member"] = user
				return
			}
			id, err := strconv.ParseInt(referralCode[4:], 10, 64)
			if err != nil{
				flash.Error("Invalid referral Code")
				this.Data["Member"] = user
				return
			}
			if ok := o.QueryTable(new(models.Member)).Filter("id", id).Exist(); !ok{
				flash.Error("Invalid referral Code")
				this.Data["Member"] = user
				return
			}
			user.ReferralId = id

		}

		user.HashPassword = user.Password// string(hashedPassword)
		//user.Password = ""
		user.Status = models.StatusActive
		user.CreatedAt = time.Now()

		bank := &models.Bank{Id:bankId}

		user.Bank = bank
		id, err := o.Insert(&user)
		if err != nil{
			o.Rollback()
			fmt.Printf("NUM: ERR: %v\n", err)
			flash.Error("Error in creating account. Please try again later")
			this.Data["Member"] = user
			return
		}

		// pair new user to make payment
		service := services.Peer2PeerService{O: o}
		err = service.CreatePayment(id)

		if err != nil{
			fmt.Sprintf("%v\n", err)
			o.Rollback()
			flash.Error("Error in creating account. Please try again later")
			this.Data["Member"] = user
			return
		}

		o.Commit()
		this.SetSession("uid", user.Id)
		beego.BeeLogger.Error("Inserted id, " + string(id))
		this.Ctx.Redirect(302, "/dashboard")
	}

}

func (this *AuthController) ForgotPassword() {

	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	if this.Ctx.Input.Method() == "POST"{
		o := orm.NewOrm()
		email := this.GetString("email")
		var member models.Member
		err := o.QueryTable(new(models.Member)).Filter("emial", email).One(&member)
		if err != nil{
			flash.Error("No record found for the email address you provided")
			this.Redirect("login", 302)
			return
		}
	}
}

func buildBankOptions(selectedBankId int64, banks []models.Bank) []dto.Option {
	var options = make([]dto.Option, len(banks))
	for index, bank := range banks{
		option := dto.Option{Id:bank.Id, Text:bank.Name}
		if selectedBankId == bank.Id{
			option.Selected = true
		}
		options[index] = option
	}
	return options
}
