package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
  username := this.GetSession("username")
  if username == nil {
    this.Ctx.Redirect(302, "/sessions/new")
  } else {
    this.Ctx.Redirect(302, "/profile")
  }
}
