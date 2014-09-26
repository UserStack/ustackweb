package controllers

import (
	"github.com/astaxie/beego"
	"github.com/UserStack/ustackweb/models"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) CreateAdmin() {
	users := models.Users().All()
	if len(users) == 0 {
		models.Users().Create("admin", "admin")
	}
	this.Redirect(beego.UrlFor("SessionsController.New"), 302)
}
