package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
}

func (this *HomeController) Get() {
	this.Redirect(beego.UrlFor("ProfileController.Get"), 302)
}
