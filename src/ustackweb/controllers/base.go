package controllers

import (
  "github.com/astaxie/beego"
)

type BaseController struct {
  beego.Controller
}

func (this *BaseController) Prepare() {
  username := this.Ctx.Input.Session("username")
  if username != nil {
    this.Data["loggedIn"] = true
    this.Data["username"] = username
  } else {
    flash := beego.NewFlash()
    flash.Error("Not logged in!")
    flash.Store(&this.Controller)
    this.Redirect("/sessions/new", 302)
    return
  }
}
