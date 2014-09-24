package controllers

type Registration struct {
    Username string
    Password string
}

type RegistrationsController struct {
  BaseController
}

func (this *RegistrationsController) Prepare() {
  this.PrepareXsrf()
  this.PrepareLayout()
  this.Layout = "layouts/default.tpl.html"
}

func (this *RegistrationsController) New() {
  this.TplNames = "registrations/new.tpl.html"
}

func (this *RegistrationsController) Create() {
  registration := Registration{}
  err := this.ParseForm(&registration)
  if err == nil && registration.Username != "foo" && registration.Username != "admin" {
    this.SetSession("username", registration.Username)
    this.RequireAuth()
    this.Redirect("/profile", 302)
  } else {
    this.Redirect("/register", 302)
  }
}
