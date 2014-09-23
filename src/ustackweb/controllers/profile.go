package controllers

import (
  "github.com/astaxie/beego"
)

type ProfileController struct {
  beego.Controller
}

func (this *ProfileController) Get() {
  this.Data["loggedIn"] = true
  this.Data["name"] = this.GetSession("username")
  this.Layout = "layouts/default.tpl.html"
  this.TplNames = "profile/index.tpl.html"
}
