package controllers

type ProfileController struct {
  BaseController
}

func (this *ProfileController) Get() {
  this.Layout = "layouts/default.tpl.html"
  this.TplNames = "profile/index.tpl.html"
}
