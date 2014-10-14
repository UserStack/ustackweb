package controllers

import (
	"strconv"
	"time"

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
	data := make(map[string]interface{})
	for _, key := range keys {
		value, err := user.Data(key)
		if err == nil {
			if key == "currentlogin" || key == "lastlogin" {
				i, err := strconv.ParseInt(value, 10, 0)
				if err == nil {
					data[key] = time.Unix(i, 0)
				}
			} else {
				data[key] = value
			}
		}
	}
	this.Data["userData"] = data
}
