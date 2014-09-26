package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
}

func (this *HomeController) Get() {
	this.Layout = "layouts/default.html.tpl"
	this.TplNames = "home/index.html.tpl"
}
