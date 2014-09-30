package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
)

type InstallController struct {
	BaseController
}

type PermissionRequirement struct {
	Name     string
	Exists   bool
	Assigned bool
}

func (this *InstallController) Index() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "config/index.html.tpl"
	rootUser, err := models.Users().FindByName("admin")
	this.Data["rootUserError"] = err
	this.Data["rootUser"] = rootUser
	permissionRequirements := this.permissionRequirements()
	groups, err := models.Groups().All()
	for _, permissionRequirement := range permissionRequirements {
		for _, group := range groups {
			if group.Name == permissionRequirement.Name {
				permissionRequirement.Exists = true
				break
			}
		}

	}
	this.Data["permissionRequirements"] = permissionRequirements
	this.Data["groupsError"] = err
}

func (this *InstallController) CreateRootUser() {
	models.Users().Create("admin", "admin")
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) CreatePermissions() {
	permissionRequirements := this.permissionRequirements()
	for _, permissionRequirement := range permissionRequirements {
		models.Groups().Create(permissionRequirement.Name)
	}
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) AssignPermissions() {
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) permissionRequirements() (permissionRequirements []*PermissionRequirement) {
	permissionRequirements = []*PermissionRequirement{
		&PermissionRequirement{Name: "perm.user.list"},
		&PermissionRequirement{Name: "perm.user.read"},
		&PermissionRequirement{Name: "perm.user.write"},
	}
	return
}
