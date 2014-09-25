package controllers

import (
	"github.com/UserStack/ustackd/backends"
	"github.com/astaxie/beego"
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

func (this *UsersController) allUsers() []backends.User {
	return []backends.User{backends.User{Uid: 1, Email: "foo"},
		backends.User{Uid: 2, Email: "admin"},
		backends.User{Uid: 3, Email: "abc"},
		backends.User{Uid: 4, Email: "def"},
		backends.User{Uid: 5, Email: "hij"},
		backends.User{Uid: 6, Email: "glk"},
		backends.User{Uid: 7, Email: "uvw"},
		backends.User{Uid: 8, Email: "xyz"}}
}

func (this *UsersController) findUser(uid int) backends.User {
	for _, user := range this.allUsers() {
		if user.Uid == uid {
			return user
		}
	}
	return backends.User{}
}

func (this *UsersController) Index() {
	this.TplNames = "users/index.tpl.html"
	users := this.allUsers()
	paginator := utils.NewPaginator(this.Ctx.Request, 3, len(users))
	this.Data["paginator"] = paginator
	this.Data["users"] = users[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func (this *UsersController) Edit() {
	this.TplNames = "users/edit.tpl.html"
	id, _ := this.GetInt(":id")
	this.Data["user"] = this.findUser(int(id))
}

func (this *UsersController) Update() {
	id, _ := this.GetInt(":id")
	user := this.findUser(int(id))
	flash := beego.NewFlash()
	flash.Notice("Updated user " + user.Email)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}

func (this *UsersController) Destroy() {
	id, _ := this.GetInt(":id")
	user := this.findUser(int(id))
	flash := beego.NewFlash()
	flash.Notice("Deleted user " + user.Email)
	flash.Store(&this.Controller)
	this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}
