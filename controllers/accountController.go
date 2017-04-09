package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/services"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"net/http"
)

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Index() {
	o := orm.NewOrm()
	accountService := services.AccountService{O:o, Ctx:this.Ctx}

	this.Data["Title"] = "My Profile"
	this.Data["CurrentMember"], _ = accountService.GetCurrentMember()

	this.Data["ProfileActive"] = "active"
	this.Layout = "shared/layout.html"
	this.TplName = "account/index.html"
}

func (this *AccountController) ChangePassword() {
	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	o := orm.NewOrm()
	accountService := services.AccountService{O:o, Ctx:this.Ctx}
	member, _ := accountService.GetCurrentMember()

	currentPassword := this.GetString("CurrentPassword")

	password := this.GetString("NewPassword")
	confirmPassword := this.GetString("ConfirmPassword")

	if password != confirmPassword{
		flash.Error("Password must be same as Confirm Password")
		this.Redirect("/account", http.StatusTemporaryRedirect)
		fmt.Println("Password must be same as Confirm Password")
		return
	}

	// validate current paasword
	if err := bcrypt.CompareHashAndPassword([]byte(member.HashPassword), []byte(currentPassword)); err != nil{
		flash.Error("Please enter your current password")
		fmt.Println("Please enter your current password")
		this.Redirect("/account", http.StatusTemporaryRedirect)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		fmt.Printf("NUM: ERR: %v\n", err)
	}

	// update password
	member.HashPassword = string(hashedPassword)
	o.Update(&member)
	flash.Success("Password changed")
	this.Redirect("/account", http.StatusTemporaryRedirect)
}
