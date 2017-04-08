package main

import (
	_ "github.com/ademuanthony/h2gh/routers"
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/context"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		allowedPaths := []string{"/auth", "/public", "/register", "/login"}
		if string(ctx.Input.URL()) ==  "/"{
			return
		}
		for _,url := range allowedPaths{
			if strings.HasPrefix(ctx.Input.URL(), url)  {
				return
			}
		}

		id := ctx.Input.Session("uid")
		_, ok := id.(int64)
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}

func main() {
	beego.Run()
}
