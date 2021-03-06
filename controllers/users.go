package controllers

import (
	"fmt"
	"github.com/UserStack/ustackweb/forms"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type UsersController struct {
	BaseController
	User                        *models.User
	UserGroups                  []models.Group
	AllGroups                   []models.Group
	AllGroupsWithoutPermissions []models.Group
	UserPermissions             []*models.UserPermission
}

type UsersFilter struct {
	GroupId string
}

func (this *UsersController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *UsersController) Index() {
	this.RequirePermissions([]string{"list_users"})
	if !this.loadAllGroups() {
		return
	}
	this.TplNames = "users/index.html.tpl"
	usersFilter := UsersFilter{GroupId: this.GetString(":groupId")}
	this.Data["usersFilter"] = usersFilter
	var users []models.User
	if usersFilter.GroupId == "" {
		users, _ = models.Users().All()
	} else {
		users, _ = models.Users().AllByGroup(usersFilter.GroupId)
	}
	paginator := pagination.SetPaginator(this.Ctx, 25, int64(len(users)))
	this.Data["users"] = users[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func (this *UsersController) New() {
	this.RequirePermissions([]string{"create_users"})
	this.TplNames = "users/new.html.tpl"
	form := forms.NewUser{}
	this.SetFormSets(&form)
}

func (this *UsersController) Create() {
	this.RequirePermissions([]string{"create_users"})
	this.TplNames = "users/new.html.tpl"
	form := forms.NewUser{}
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}
	created, id, _ := models.Users().Create(form.Username, form.Password)
	if created {
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", fmt.Sprintf("%d", id)), 302)
	} else { // backend error
		flash := beego.NewFlash()
		flash.Error("Could not create user!")
		flash.Store(&this.Controller)
		this.TplNames = "users/new.html.tpl"
	}
}

func (this *UsersController) Edit() {
	this.RequirePermissions([]string{"read_users"})
	if !this.loadUser() || !this.loadUserGroups() || !this.loadUserPermissions() {
		return
	}
	this.TplNames = "users/edit.html.tpl"
}

type GroupMembership struct {
	Group    models.Group
	IsMember bool
}

func (this *UsersController) EditGroups() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() || !this.loadUserGroups() || !this.loadAllGroupsWithoutPermissions() {
		return
	}
	groupMemberships := make([]GroupMembership, len(this.AllGroupsWithoutPermissions))
	for idx, group := range this.AllGroupsWithoutPermissions {
		membership := GroupMembership{Group: group}
		for _, group2 := range this.UserGroups {
			if group == group2 {
				membership.IsMember = true
				break
			}
		}
		groupMemberships[idx] = membership
	}
	this.Data["groupMemberships"] = groupMemberships
	this.TplNames = "users/edit_groups.html.tpl"
}

func (this *UsersController) EditPermissions() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() || !this.loadUserGroups() || !this.loadUserPermissions() {
		return
	}
	this.TplNames = "users/edit_permissions.html.tpl"
}

func (this *UsersController) AddUserToGroup() {
	this.RequirePermissions([]string{"update_users"})
	if models.Permissions().IsPermissionGroupName(this.GetString(":groupId")) {
		this.RequirePermissions([]string{"grant_permissions"})
	}
	if !this.loadUser() {
		return
	}
	models.Users().AddUserToGroup(this.GetString(":id"), this.GetString(":groupId"))
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) RemoveUserFromGroup() {
	this.RequirePermissions([]string{"update_users"})
	if models.Permissions().IsPermissionGroupName(this.GetString(":groupId")) {
		this.RequirePermissions([]string{"revoke_permissions"})
	}
	if !this.loadUser() {
		return
	}
	models.Users().RemoveUserFromGroup(this.GetString(":id"), this.GetString(":groupId"))
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) EditUsername() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit_username.html.tpl"
	form := forms.EditUsername{}
	this.SetFormSets(&form)
}

func (this *UsersController) UpdateUsername() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit_username.html.tpl"
	form := forms.EditUsername{}
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}

	updated, backendErr := models.Users().UpdateUsername(this.GetString(":id"), form.ConfirmPassword, form.Username)
	if updated { // success
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
	} else { // backend error
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Could not update user! %s", backendErr))
		flash.Store(&this.Controller)
	}
}

func (this *UsersController) EditPassword() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit_password.html.tpl"
	form := forms.EditPassword{}
	this.SetFormSets(&form)
}

func (this *UsersController) UpdatePassword() {
	this.RequirePermissions([]string{"update_users"})
	if !this.loadUser() {
		return
	}

	this.TplNames = "users/edit_password.html.tpl"
	form := forms.EditPassword{}
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}

	updated, backendErr := models.Users().UpdatePassword(this.GetString(":id"), form.OldPassword, form.NewPassword)
	if updated { // success
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
	} else { // backend error
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Could not update user! %s", backendErr))
		flash.Store(&this.Controller)
	}
}

func (this *UsersController) Destroy() {
	this.RequirePermissions([]string{"delete_users"})
	id, _ := this.GetInt64(":id")
	user, _ := models.Users().Find(id)
	models.Users().Destroy(user.Name)
	flash := beego.NewFlash()
	flash.Notice("Deleted user " + user.Name)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}

func (this *UsersController) Enable() {
	this.RequirePermissions([]string{"enable_users"})
	id, _ := this.GetInt64(":id")
	user, _ := models.Users().Find(id)
	models.Users().Enable(user.Name)
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) Disable() {
	this.RequirePermissions([]string{"disable_users"})
	id, _ := this.GetInt64(":id")
	user, _ := models.Users().Find(id)
	models.Users().Disable(user.Name)
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) loadUser() (loaded bool) {
	intId, _ := this.GetInt64(":id")
	user, err := models.Users().Find(intId)
	loaded = err == nil
	if !loaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	this.User = user
	this.Data["user"] = user
	return
}

func (this *UsersController) loadUserGroups() (loaded bool) {
	groups, err := models.Groups().AllByUserWithoutPermissions(this.GetString(":id"))
	loaded = err == nil
	if !loaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	this.UserGroups = groups
	this.Data["userGroups"] = groups
	return
}

func (this *UsersController) loadUserPermissions() bool {
	userPermissions := models.UserPermissions().All(this.GetString(":id"))
	this.UserPermissions = userPermissions
	this.Data["userPermissions"] = userPermissions
	return true
}

func (this *UsersController) loadAllGroups() (loaded bool) {
	groups, err := models.Groups().All()
	loaded = err == nil
	if !loaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	this.AllGroups = groups
	this.Data["allGroups"] = groups
	return
}

func (this *UsersController) loadAllGroupsWithoutPermissions() (loaded bool) {
	groups, err := models.Groups().AllWithoutPermissions()
	loaded = err == nil
	if !loaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	this.AllGroupsWithoutPermissions = groups
	this.Data["allGroupsWithoutPermissions"] = groups
	return
}
