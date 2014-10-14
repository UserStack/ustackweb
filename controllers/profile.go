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
	user, err := models.Users().FindByName("admin")
	if err != nil {
		return
	}
	keys, err := user.DataKeys()
	if err != nil {
		return
	}
	data := make(map[string]string)
	for _, key := range keys {
		value, err := user.Data(key)
		if err == nil {
			data[key] = value
		}
	}
	this.Data["userData"] = data
}
