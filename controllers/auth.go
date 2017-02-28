package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/TaxMaster/models"
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
		user := models.UserAuth{}

		err := o.QueryTable("user_auth").Filter("email", email).One(&user)
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


		if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password)); err == nil{
			this.SetSession("uid", user.Id)
			this.Redirect("/", 302)
			return
		}
		flash.Error("Invalid credetial p")
		flash.Store(&this.Controller)
	}
}

func (this *AuthController) Logout() {
	this.DelSession("uid")
	this.Redirect("/auth/login", 308)//removed
}

func (this *AuthController) Index()  {

}

func (this *AuthController) Create() {
	if this.Ctx.Input.Method() == "POST"{
		user := models.UserAuth{}
		if err := this.ParseForm(&user); err != nil {
			//handle error
			fmt.Printf("NUM: ERR: %v\n", err)
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("NUM: ERR: %v\n", err)
		}
		user.HashPassword = string(hashedPassword)
		user.Password = ""
		o := orm.NewOrm()
		id, err := o.Insert(&user)
		if err != nil{
			fmt.Printf("NUM: ERR: %v\n", err)
		}
		beego.BeeLogger.Error("Inserted id, " + string(id))
		this.Ctx.Redirect(302, "/accounts")
	}
}