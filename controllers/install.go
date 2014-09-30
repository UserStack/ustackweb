package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
)

type InstallController struct {
	BaseController
}

type GroupRequirenment struct {
	Name   string
	Exists bool
}

func (this *InstallController) Index() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "config/index.html.tpl"
	rootUser, err := models.Users().FindByName("admin")
	this.Data["rootUserError"] = err
	this.Data["rootUser"] = rootUser
	groupRequirements := []*GroupRequirenment{
		&GroupRequirenment{Name: "perm.user.list"},
		&GroupRequirenment{Name: "perm.user.read"},
		&GroupRequirenment{Name: "perm.user.write"},
	}
	groups, err := models.Groups().All()
	for _, groupRequirement := range groupRequirements {
		for _, group := range groups {
			if group.Name == groupRequirement.Name {
				groupRequirement.Exists = true
				break
			}
		}
	}
	this.Data["groupRequirements"] = groupRequirements
	this.Data["groupsError"] = err
}

func (this *InstallController) CreateRootUser() {
	models.Users().Create("admin", "admin")
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}
