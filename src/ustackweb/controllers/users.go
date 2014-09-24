package controllers

type UsersController struct {
  BaseController
}

func (this *UsersController) Prepare() {
  this.PrepareXsrf()
  this.RequireAuth()
  this.PrepareLayout()
  this.Layout = "layouts/default.tpl.html"
}

func (this *UsersController) Get() {
  this.TplNames = "users/index.tpl.html"
  this.Data["users"] = []string{"foo","admin"}
}
