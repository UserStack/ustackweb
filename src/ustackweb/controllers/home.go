package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
  this.PrepareXsrf();
  this.RequireAuth();
}

func (this *HomeController) Get() {
  this.Redirect("/profile", 302)
}
