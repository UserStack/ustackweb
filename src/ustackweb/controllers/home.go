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
  this.Redirect("/profile", 302)
}
