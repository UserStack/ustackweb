package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"ustackweb/models"
	"ustackweb/utils"
)

type UserForm struct {
	Username string
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
	paginator := utils.NewPaginator(this.Ctx.Request, 3, len(users))
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
			flash.Error("Insufficient data!")
			flash.Store(&this.Controller)
			this.TplNames = "users/new.html.tpl"
		} else {
			this.Redirect(beego.UrlFor("UsersController.Index"), 302)
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

func (this *UsersController) Update() {
	id, _ := this.GetInt(":id")
	user := models.Users().Find(id)
	flash := beego.NewFlash()
	flash.Notice("Updated user " + user.Email)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}

func (this *UsersController) Destroy() {
	id, _ := this.GetInt(":id")
	user := models.Users().Find(id)
	flash := beego.NewFlash()
	flash.Notice("Deleted user " + user.Email)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}
