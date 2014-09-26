package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	wetalkutils "github.com/beego/wetalk/modules/utils"
	"ustackweb/models"
	"ustackweb/utils"
)

type UserForm struct {
	Username    string
	Password    string
	OldPassword string
}

type UsersController struct {
	BaseController
}

func (this *UsersController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *UsersController) Index() {
	this.TplNames = "users/index.html.tpl"
	users := models.Users().All()
	paginator := wetalkutils.NewPaginator(this.Ctx.Request, 25, len(users))
	this.Data["paginator"] = paginator
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
			created, id := models.Users().Create(userForm.Username, userForm.Password)
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
	this.Data["user"] = models.Users().Find(id)
}

func (this *UsersController) UpdateUsername() {
	id := this.GetString(":id")
	intId, _ := this.GetInt(":id")
	this.Data["user"] = models.Users().Find(intId)
	userForm := UserForm{}
	err := this.ParseForm(&userForm)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(userForm.Username, "Username")
		valid.Required(userForm.Password, "Password")
		if valid.HasErrors() {
			this.Data["UpdateUsernameErrors"] = valid.Errors
			flash := beego.NewFlash()
			flash.Error("Could not change username.")
			flash.Store(&this.Controller)
			this.TplNames = "users/edit.html.tpl"
		} else {
			updated := models.Users().UpdateUsername(id, userForm.Password, userForm.Username)
			if updated {
				this.Redirect(beego.UrlFor("UsersController.Edit", ":id", string(id)), 302)
			} else {
				flash := beego.NewFlash()
				flash.Error("Could not update user!")
				flash.Store(&this.Controller)
				this.TplNames = "users/edit.html.tpl"
			}
		}
	} else {
		this.TplNames = "users/edit.html.tpl"
	}
}

func (this *UsersController) UpdatePassword() {
	id := this.GetString(":id")
	intId, _ := this.GetInt(":id")
	this.Data["user"] = models.Users().Find(intId)
	userForm := UserForm{}
	err := this.ParseForm(&userForm)
	if err == nil {
		valid := validation.Validation{}
		valid.Required(userForm.Password, "Password")
		valid.Required(userForm.OldPassword, "OldPassword")
		if valid.HasErrors() {
			this.Data["UpdatePasswordErrors"] = valid.Errors
			flash := beego.NewFlash()
			flash.Error("Could not change password.")
			flash.Store(&this.Controller)
			this.TplNames = "users/edit.html.tpl"
		} else {
			updated := models.Users().UpdatePassword(id, userForm.OldPassword, userForm.Password)
			if updated {
				this.Redirect(beego.UrlFor("UsersController.Edit", ":id", string(id)), 302)
			} else {
				flash := beego.NewFlash()
				flash.Error("Could not update user!")
				flash.Store(&this.Controller)
				this.TplNames = "users/edit.html.tpl"
			}
		}
	} else {
		this.TplNames = "users/edit.html.tpl"
	}
}

func (this *UsersController) Destroy() {
	id, _ := this.GetInt(":id")
	user := models.Users().Find(id)
	models.Users().Destroy(user.Name)
	flash := beego.NewFlash()
	flash.Notice("Deleted user " + user.Name)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}
