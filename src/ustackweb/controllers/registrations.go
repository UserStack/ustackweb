package controllers

import (
	"github.com/astaxie/beego"
	"ustackweb/models"
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
	this.Layout = "layouts/default.html.tpl"
}

func (this *RegistrationsController) New() {
	this.TplNames = "registrations/new.html.tpl"
}

func (this *RegistrationsController) Create() {
	registration := Registration{}
	err := this.ParseForm(&registration)
	if err == nil {
		models.Users().Create(registration.Username, registration.Password)
		this.Redirect(beego.UrlFor("SessionsController.New"), 302)
	} else {
		this.Redirect(beego.UrlFor("RegistrationsController.New"), 302)
	}
}
