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

func (this *InstallController) rootUserId() string {
	return "admin"
}

func (this *InstallController) permissionRequirements() (permissionRequirements []*PermissionRequirement) {
	allPermissions := models.Permissions().AllGroupNames()
	permissionRequirements = make([]*PermissionRequirement, len(allPermissions))
	for idx, groupName := range allPermissions {
		permissionRequirements[idx] = &PermissionRequirement{Name: groupName}
	}
	return
}

func (this *InstallController) Index() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "config/index.html.tpl"
	rootUser, err := models.Users().FindByName(this.rootUserId())
	this.Data["rootUserError"] = err
	this.Data["rootUser"] = rootUser
	this.Data["hasRootUser"] = rootUser != nil
	this.Data["hasRootUserError"] = err != nil
	groups, err := models.Groups().All()
	this.Data["groupsError"] = err
	userGroups, err := models.Groups().AllByUser(this.rootUserId())
	this.Data["userGroupsError"] = err
	permissionRequirements := this.permissionRequirements()
	for _, permissionRequirement := range permissionRequirements {
		for _, group := range groups {
			if group.Name == permissionRequirement.Name {
				permissionRequirement.Exists = true
				break
			}
		}
		for _, userGroup := range userGroups {
			if userGroup.Name == permissionRequirement.Name {
				permissionRequirement.Assigned = true
				break
			}
		}
	}
	this.Data["permissionRequirements"] = permissionRequirements
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
	permissionRequirements := this.permissionRequirements()
	for _, permissionRequirement := range permissionRequirements {
		models.Users().AddUserToGroup(this.rootUserId(), permissionRequirement.Name)
	}
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}
