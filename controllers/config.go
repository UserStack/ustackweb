package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
)

type ConfigController struct {
	BaseController
}

func (this *ConfigController) CreateAdmin() {
	users, _ := models.Users().All()
	if len(users) == 0 {
		models.Users().Create("admin", "admin")
	}
	this.Redirect(beego.UrlFor("SessionsController.New"), 302)
}
