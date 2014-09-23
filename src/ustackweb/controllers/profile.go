package controllers

import (
  "github.com/astaxie/beego"
)

type ProfileController struct {
  beego.Controller
}

func (this *ProfileController) Get() {
  this.Data["name"] = this.GetSession("username")
  this.Layout = "layouts/default.tpl"
  this.TplNames = "profile/index.tpl"
}
