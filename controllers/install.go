package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
)

type InstallController struct {
	BaseController
}

func (this *InstallController) Index() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "config/index.html.tpl"
	rootUser, err := models.Users().FindByName("admin")
	this.Data["rootUserError"] = err
	this.Data["rootUser"] = rootUser
	groups, err := models.Groups().All()
	this.Data["groupsError"] = err
	this.Data["groups"] = groups
}

func (this *InstallController) CreateRootUser() {
	models.Users().Create("admin", "admin")
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}
