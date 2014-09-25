package controllers

import (
	"github.com/astaxie/beego"
)

type Registration struct {
	Username string
	Password string
}

type RegistrationsController struct {
	BaseController
}

func (this *RegistrationsController) Prepare() {
	this.PrepareXsrf()
	this.PrepareLayout()
	this.Layout = "layouts/default.tpl.html"
}

func (this *RegistrationsController) New() {
	this.TplNames = "registrations/new.tpl.html"
}

func (this *RegistrationsController) Create() {
	registration := Registration{}
	err := this.ParseForm(&registration)
	if err == nil && registration.Username != "foo" && registration.Username != "admin" {
		this.SetSession("username", registration.Username)
		this.RequireAuth()
		this.Redirect(beego.UrlFor("ProfileController.Get"), 302)
	} else {
		this.Redirect(beego.UrlFor("RegistrationsController.New"), 302)
	}
}
