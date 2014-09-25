package controllers

import (
	"github.com/astaxie/beego"
	"ustackweb/models"
	"ustackweb/utils"
)

type UsersController struct {
	BaseController
}

func (this *UsersController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.tpl.html"
}

func (this *UsersController) Index() {
	this.TplNames = "users/index.tpl.html"
	users := models.Users().All()
	paginator := utils.NewPaginator(this.Ctx.Request, 3, len(users))
	this.Data["paginator"] = paginator
	this.Data["users"] = users[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func (this *UsersController) New() {
	this.TplNames = "users/new.tpl.html"
}

func (this *UsersController) Create() {
	flash := beego.NewFlash()
	flash.Notice("Updated user " + this.GetString("username"))
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}

func (this *UsersController) Edit() {
	this.TplNames = "users/edit.tpl.html"
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
