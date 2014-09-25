package controllers

import (
  "ustackweb/utils"
  "github.com/UserStack/ustackd/backends"
  "github.com/astaxie/beego"
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
  users := []backends.User{backends.User{Uid:1,Email:"foo"},
                           backends.User{Uid:2,Email:"admin"},
                           backends.User{Uid:3,Email:"abc"},
                           backends.User{Uid:4,Email:"def"},
                           backends.User{Uid:5,Email:"hij"},
                           backends.User{Uid:6,Email:"glk"},
                           backends.User{Uid:7,Email:"uvw"},
                           backends.User{Uid:8,Email:"xyz"}}
  paginator := utils.NewPaginator(this.Ctx.Request, 3, len(users))
  this.Data["paginator"] = paginator
  this.Data["users"] = users[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(users))]
}

func (this *UsersController) Edit() {
  // this.GetString(":id")
  this.TplNames = "users/edit.tpl.html"
  this.Data["user"] = backends.User{Uid:1,Email:"foo"}
}

func (this *UsersController) Update() {
  // this.GetString(":id")
  flash := beego.NewFlash()
  flash.Notice("Updated")
  flash.Store(&this.Controller)
  this.Redirect(beego.UrlFor("UsersController.Index"), 302)
}
