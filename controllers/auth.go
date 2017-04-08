package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
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

		err := o.QueryTable("member").Filter("email", email).One(&user)
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


		if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password)); err != nil{
			this.SetSession("uid", user.Id)
			fmt.Printf("User Id: %v\n", user.Id)
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

	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	var banks []models.Bank
	_, err := o.QueryTable(new(models.Bank)).All(&banks)
	if err != nil{
		panic(err)
	}

	this.Data["Banks"] = banks
	this.Data["Title"] = "Create Account"

	this.TplName = "auth/register.html";

	if this.Ctx.Input.Method() == "POST"{
		user := models.Member{}
		if err := this.ParseForm(&user); err != nil {
			//handle error
			fmt.Printf("NUM: ERR: %v\n", err)
		}

		if this.GetString("ConfirmPassword") != user.Password{
			fmt.Printf("NUM: ERR: %v\n", err)
			flash.Error("Password and Confirm password must be the same")
			this.Data["Member"] = user
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("NUM: ERR: %v\n", err)
		}
		user.HashPassword = string(hashedPassword)
		user.Password = ""
		user.Status = models.StatusActive

		bankId, _ := this.GetInt64("BankId")
		bank := &models.Bank{Id:bankId}

		user.Bank = bank
		id, err := o.Insert(&user)
		if err != nil{
			fmt.Printf("NUM: ERR: %v\n", err)
			flash.Error("Error in creating account. Please try again later")
			this.Data["Member"] = user
			return
		}

		this.SetSession("uid", user.Id)
		beego.BeeLogger.Error("Inserted id, " + string(id))
		this.Ctx.Redirect(302, "/dashboard")
	}


}