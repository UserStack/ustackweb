package main

import (
	_ "ustackweb/routers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/context"
)

var FilterAuthentication = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("username").(string)
    if !ok {
      ctx.Redirect(302, "/sessions/new")
    }
}

func main() {
  beego.SessionOn = true
  beego.EnableAdmin = true
  beego.InsertFilter("/profile", beego.BeforeRouter, FilterAuthentication)
	beego.Run()
}

