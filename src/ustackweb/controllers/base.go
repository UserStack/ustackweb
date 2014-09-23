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
    this.Redirect("/sessions/new", 302)
    return
  }
}
