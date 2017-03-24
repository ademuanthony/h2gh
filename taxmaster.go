package main

import (
	_ "github.com/ademuanthony/TaxMaster/routers"
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/context"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		allowedPaths := []string{"/auth", "/public"}
		if string(ctx.Input.URL()) ==  "/"{
			return
		}
		for _,url := range allowedPaths{
			if strings.HasPrefix(ctx.Input.URL(), url)  {
				return
			}
		}

		_, ok := ctx.Input.Session("uid").(int)
		if !ok {
			ctx.Redirect(302, "/auth/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)


	beego.Run()
}

func main() {
	beego.Run()
}
