package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
  this.RequireAuth();
}

func (this *HomeController) Get() {
  this.Ctx.Redirect(302, "/profile")
}
