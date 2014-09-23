package controllers

import (
  "github.com/astaxie/beego"
)

type SessionsController struct {
  beego.Controller
}

type Login struct {
    Username string `form:"username"`
}

func (this *SessionsController) New() {
  this.Data["Form"] = &Login{}
  this.Layout = "layouts/default.tpl.html"
  this.TplNames = "sessions/new.tpl.html"
}

func (this *SessionsController) Create() {
  login := Login{}
  if err := this.ParseForm(&login); err != nil {
    this.Ctx.Redirect(302, "/sessions/new")
  } else {
    this.SetSession("username", login.Username)
    this.Ctx.Redirect(302, "/profile")
  }
}

func (this *SessionsController) Destroy() {
  this.DelSession("username")
  this.Ctx.Redirect(302, "/sessions/new")
}
