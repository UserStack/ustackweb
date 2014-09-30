package controllers

import (
	"fmt"
	"github.com/UserStack/ustackweb/forms"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserForm struct {
	Username    string
	Password    string
	OldPassword string
}

type UsersController struct {
	BaseController
	User *models.User
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
}

func (this *UsersController) Create() {
	userForm := UserForm{}
	err := this.ParseForm(&userForm)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(userForm.Username, "Username")
		if valid.HasErrors() {
			this.Data["Errors"] = valid.Errors
			flash := beego.NewFlash()
			flash.Error("Could not create user.")
			flash.Store(&this.Controller)
			this.TplNames = "users/new.html.tpl"
		} else {
			created, id, _ := models.Users().Create(userForm.Username, userForm.Password)
			if created {
				this.Redirect(beego.UrlFor("UsersController.Edit", ":id", fmt.Sprintf("%d", id)), 302)
			} else {
				flash := beego.NewFlash()
				flash.Error("Could not create user!")
				flash.Store(&this.Controller)
				this.TplNames = "users/new.html.tpl"
			}
		}
	} else {
		this.TplNames = "users/new.html.tpl"
	}
}

func (this *UsersController) Edit() {
	this.TplNames = "users/edit.html.tpl"
	id, _ := this.GetInt(":id")
	user, error := models.Users().Find(id)
	if error == nil {
		this.Data["user"] = user
	} else {
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
}

func (this *UsersController) EditUsername() {
	this.TplNames = "users/username.html.tpl"
	id, _ := this.GetInt(":id")
	user, error := models.Users().Find(id)
	if error == nil {
		this.Data["user"] = user
		this.Data["UsernameForm"] = forms.EditUsername{XsrfHtml: this.XsrfFormHtml(), User: user}
	} else {
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
}

func (this *UsersController) EditPassword() {
	this.TplNames = "users/password.html.tpl"
	id, _ := this.GetInt(":id")
	user, error := models.Users().Find(id)
	if error == nil {
		this.Data["user"] = user
		this.Data["PasswordForm"] = forms.EditPassword{XsrfHtml: this.XsrfFormHtml(), User: user}
	} else {
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
}

func (this *UsersController) loadUser() (userLoaded bool) {
	intId, _ := this.GetInt(":id")
	user, err := models.Users().Find(intId)
	this.User = user
	userLoaded = err == nil
	if !userLoaded { // user not found
		this.Redirect(beego.UrlFor("UsersController.Index"), 302)
	}
	return
}

func (this *UsersController) parseUserForm() (userForm *UserForm) {
	userForm = &UserForm{}
	err := this.ParseForm(&userForm)
	if err != nil { // form broken / hacked
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
		return
	}
	return
}

func (this *UsersController) UpdateUsername() {
	if !this.loadUser() {
		return
	}

	userForm := this.parseUserForm()
	if userForm == nil {
		return
	}

	valid := validation.Validation{}
	valid.Required(userForm.Username, "Username")
	valid.Required(userForm.Password, "Password")
	if valid.HasErrors() { // validation errors
		usernameForm := forms.EditUsername{XsrfHtml: this.XsrfFormHtml(), User: this.User, ValidationErrors: valid.Errors}
		this.Data["user"] = this.User
		this.Data["UsernameForm"] = usernameForm
		flash := beego.NewFlash()
		flash.Error("Could not change username.")
		flash.Store(&this.Controller)
		this.TplNames = "users/username.html.tpl"
		return
	}

	updated, backendErr := models.Users().UpdateUsername(this.GetString(":id"), userForm.Password, userForm.Username)
	if updated { // success
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
	} else {
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Could not update user! %s", backendErr))
		flash.Store(&this.Controller)
		this.TplNames = "users/username.html.tpl"
	}
}

func (this *UsersController) UpdatePassword() {
	if !this.loadUser() {
		return
	}

	userForm := this.parseUserForm()
	if userForm == nil {
		return
	}

	valid := validation.Validation{}
	valid.Required(userForm.Password, "Password")
	valid.Required(userForm.OldPassword, "OldPassword")
	if valid.HasErrors() { // validation errors
		passwordForm := forms.EditPassword{XsrfHtml: this.XsrfFormHtml(), User: this.User, ValidationErrors: valid.Errors}
		this.Data["user"] = this.User
		this.Data["PasswordForm"] = passwordForm
		flash := beego.NewFlash()
		flash.Error("Could not change password.")
		flash.Store(&this.Controller)
		this.TplNames = "users/password.html.tpl"
		return
	}

	updated, backendErr := models.Users().UpdateUsername(this.GetString(":id"), userForm.Password, userForm.Username)
	if updated { // success
		this.Redirect(beego.UrlFor("UsersController.Edit", ":id", this.GetString(":id")), 302)
	} else {
		flash := beego.NewFlash()
		flash.Error(fmt.Sprintf("Could not update user! %s", backendErr))
		flash.Store(&this.Controller)
		this.TplNames = "users/password.html.tpl"
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
