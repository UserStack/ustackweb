package controllers

import (
	"github.com/astaxie/beego"
	"github.com/UserStack/ustackweb/utils"
)

type Permissions struct {
	Users  bool
	Groups bool
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) PrepareLayout() {
	this.Data["context"] = utils.NewContext(this.GetControllerAndAction())
	this.Layout = "layouts/default.html.tpl"
	this.Data["Lang"] = "en-US"
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
		this.Data["permissions"] = &Permissions{Users: username == "admin",
			Groups: username == "admin"}
	} else {
		this.RequireAuthFailed()
	}
}

func (this *BaseController) RequireAuthFailed() {
	flash := beego.NewFlash()
	flash.Error("Not logged in!")
	flash.Store(&this.Controller)
	this.Redirect("/sign_in", 302)
}
