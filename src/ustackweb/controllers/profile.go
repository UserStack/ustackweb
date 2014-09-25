package controllers

type ProfileController struct {
	BaseController
}

func (this *ProfileController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
}

func (this *ProfileController) Get() {
	this.Layout = "layouts/default.tpl.html"
	this.TplNames = "profile/index.tpl.html"
}
