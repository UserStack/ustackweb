package controllers

type SessionsController struct {
  BaseController
}

type Login struct {
    Username string
    Password string
}

func (this *SessionsController) Prepare() {
  this.PrepareXsrf()
  this.PrepareLayout()
}

func (this *SessionsController) New() {
  this.Data["Form"] = &Login{}
  this.TplNames = "sessions/new.tpl.html"
}

func (this *SessionsController) Create() {
  login := Login{}
  err := this.ParseForm(&login)
  if err == nil && (login.Username == "foo" || login.Username == "admin") && login.Password == "bar" {
    this.SetSession("username", login.Username)
    this.RequireAuth()
    this.Redirect("/profile", 302)
  } else {
    this.RequireAuthFailed()
  }
}

func (this *SessionsController) Destroy() {
  this.DelSession("username")
  this.Redirect("/sign_in", 302)
}
