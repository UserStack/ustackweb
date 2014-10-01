package controllers

import (
	"fmt"
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
)

type InstallController struct {
	BaseController
}

type PermissionRequirement struct {
	Permission *models.Permission
	Exists     bool
	Assigned   bool
}

func (this *InstallController) rootUserId() string {
	return "admin"
}

func (this *InstallController) permissionRequirements() (permissionRequirements []*PermissionRequirement) {
	allPermissions := models.Permissions().AllNames()
	permissionRequirements = make([]*PermissionRequirement, len(allPermissions))
	for idx, name := range allPermissions {
		permissionRequirements[idx] = &PermissionRequirement{Permission: &models.Permission{Name: name}}
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
			if group.Name == permissionRequirement.Permission.GroupName() {
				permissionRequirement.Exists = true
				break
			}
		}
		for _, userGroup := range userGroups {
			if userGroup.Name == permissionRequirement.Permission.GroupName() {
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
	models.Permissions().Create()
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) AssignPermissions() {
	names := models.Permissions().AllNames()
	for _, name := range names {
		models.Permissions().Allow(this.rootUserId(), name)
	}
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) DropDatabase() {
	users, _ := models.Users().All()
	for _, user := range users {
		models.Users().Destroy(fmt.Sprintf("%s", user.Uid))
	}
	groups, _ := models.Groups().All()
	for _, group := range groups {
		models.Groups().Destroy(group.Name)
	}
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}
