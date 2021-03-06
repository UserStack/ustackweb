package controllers

import (
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
	allPermissions := models.Permissions().AllInternal()
	permissionRequirements = make([]*PermissionRequirement, len(allPermissions))
	for idx, permission := range allPermissions {
		permissionRequirements[idx] = &PermissionRequirement{Permission: permission}
	}
	return
}

func (this *InstallController) Prepare() {
	this.PrepareLayout()
}

func (this *InstallController) Index() {
	this.TplNames = "config/index.html.tpl"
	rootUser, err := models.Users().FindByName(this.rootUserId())
	this.Data["rootUserError"] = err
	this.Data["rootUser"] = rootUser
	this.Data["hasRootUser"] = rootUser != nil
	this.Data["hasRootUserError"] = err != nil
	groups, err := models.Groups().All()
	this.Data["groupsError"] = err
	abilities := models.UserPermissions().Abilities(this.rootUserId())
	permissionRequirements := this.permissionRequirements()
	for _, permissionRequirement := range permissionRequirements {
		for _, group := range groups {
			if group.Name == permissionRequirement.Permission.GroupName() {
				permissionRequirement.Exists = true
				break
			}
		}
		permissionRequirement.Assigned = abilities[permissionRequirement.Permission.Name]
	}
	this.Data["permissionRequirements"] = permissionRequirements
}

func (this *InstallController) CreateRootUser() {
	models.Users().Create("admin", "admin")
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) CreatePermissions() {
	models.Permissions().CreateAllInternal()
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) AssignPermissions() {
	models.UserPermissions().AllowAll(this.rootUserId())
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}

func (this *InstallController) DropDatabase() {
	users, _ := models.Users().All()
	for _, user := range users {
		models.Users().Destroy(user.Name)
	}
	groups, _ := models.Groups().All()
	for _, group := range groups {
		models.Groups().Destroy(group.Name)
	}
	this.Redirect(beego.UrlFor("InstallController.Index"), 302)
}
