package controllers

import (
	"fmt"
	"github.com/UserStack/ustackweb/forms"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	wetalkutils "github.com/beego/wetalk/modules/utils"
)

type UsersController struct {
	BaseController
	User       *models.User
	UserGroups []models.Group
}

func (this *UsersController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *UsersController) Index() {
	this.TplNames = "users/index.html.tpl"
	users, _ := models.Users().All()
	paginator := this.SetPaginator(25, int64(len(users)))
	this.Data["users"] = users[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func (this *UsersController) New() {
	this.TplNames = "users/new.html.tpl"
	form := forms.NewUser{}
	this.SetFormSets(&form)
}

func (this *UsersController) Create() {
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
	if !this.loadUser() || !this.loadUserGroups() {
		return
	}
	this.TplNames = "users/edit.html.tpl"
	form := forms.AddUserToGroup{}
	form.Groups, _ = models.Groups().All()
	this.SetFormSets(&form)
}

func (this *UsersController) AddUserToGroup() {
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit.html.tpl"
	form := forms.AddUserToGroup{}
	form.Groups, _ = models.Groups().All()
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}

	updated, backendErr := models.Users().AddUserToGroup(this.GetString(":id"), wetalkutils.ToStr(form.GroupId))
	if updated { // success
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
	} else { // backend error
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Could not update user! %s", backendErr))
		flash.Store(&this.Controller)
	}
}

func (this *UsersController) RemoveUserFromGroup() {
	if !this.loadUser() {
		return
	}
	models.Users().RemoveUserFromGroup(this.GetString(":id"), this.GetString(":groupId"))
	this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
}

func (this *UsersController) EditUsername() {
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit_username.html.tpl"
	form := forms.EditUsername{}
	this.SetFormSets(&form)
}

func (this *UsersController) UpdateUsername() {
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
	if !this.loadUser() {
		return
	}
	this.TplNames = "users/edit_password.html.tpl"
	form := forms.EditPassword{}
	this.SetFormSets(&form)
}

func (this *UsersController) UpdatePassword() {
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
	id, _ := this.GetInt(":id")
	user, _ := models.Users().Find(id)
	models.Users().Destroy(user.Name)
	flash := beego.NewFlash()
	flash.Notice("Deleted user " + user.Name)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}

func (this *UsersController) Enable() {
	id, _ := this.GetInt(":id")
	user, _ := models.Users().Find(id)
	models.Users().Enable(user.Name)
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) Disable() {
	id, _ := this.GetInt(":id")
	user, _ := models.Users().Find(id)
	models.Users().Disable(user.Name)
	this.Redirect(this.Ctx.Input.Refer(), 302)
}

func (this *UsersController) loadUser() (loaded bool) {
	intId, _ := this.GetInt(":id")
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
	userGroups, err := models.Groups().AllByUser(this.GetString(":id"))
	loaded = err == nil
	if !loaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	this.UserGroups = userGroups
	this.Data["userGroups"] = userGroups
	return
}
