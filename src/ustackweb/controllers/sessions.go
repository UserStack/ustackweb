package controllers

type SessionsController struct {
  BaseController
}

type Login struct {
    Username string `form:"username"`
}

func (this *SessionsController) New() {
  this.Data["Form"] = &Login{}
  this.Layout = "layouts/default.tpl.html"
  this.TplNames = "sessions/new.tpl.html"
}

func (this *SessionsController) Create() {
  login := Login{}
  err := this.ParseForm(&login)
  if err == nil && login.Username == "foo" {
    this.SetSession("username", login.Username)
    this.RequireAuth()
    this.Redirect(302, "/profile")
  } else {
    this.RequireAuthFailed()
  }
}

func (this *SessionsController) Destroy() {
  this.DelSession("username")
  this.Ctx.Redirect(302, "/sessions/new")
}
