package controllers

import (
	"github.com/UserStack/ustackweb/models"
)

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
}

func (this *ProfileController) Get() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "profile/index.html.tpl"
	user, _ := models.Users().FindByName("admin")
	this.Data["userDataKeys"], _ = user.DataKeys()
}
