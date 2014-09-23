package controllers

import (
  "github.com/astaxie/beego"
)

type BaseController struct {
  beego.Controller
}

func (this *BaseController) PrepareXsrf() {
  this.Data["xsrf_token"] = this.XsrfToken()
  this.Data["xsrf_html"] = this.XsrfFormHtml()
}

func (this *BaseController) RequireAuth() {
  username := this.Ctx.Input.Session("username")
  if username != nil {
    this.Data["loggedIn"] = true
    this.Data["username"] = username
  } else {
    this.RequireAuthFailed()
  }
}

func (this *BaseController) RequireAuthFailed() {
  flash := beego.NewFlash()
  flash.Error("Not logged in!")
  flash.Store(&this.Controller)
  this.Redirect("/sessions/new", 302)
}
