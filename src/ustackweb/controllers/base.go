package controllers

import (
  "github.com/astaxie/beego"
  "strings"
)

type Permissions struct {
  Users bool
  Groups bool
}

type Context struct {
  controllerName string
  actionName string
}

func (this *Context) ControllerName() string {
  return strings.ToLower(strings.TrimSuffix(this.controllerName, "Controller"))
}

func (this *Context) ActionName() string {
  return strings.ToLower(this.actionName)
}

func (this *Context) Is(controllerAndAction string) bool {
  return strings.ToLower(this.controllerName + "." + this.actionName) == strings.ToLower(controllerAndAction)
}

type BaseController struct {
  beego.Controller
}

func (this *BaseController) PrepareLayout() {
  controllerName, actionName := this.GetControllerAndAction()
  this.Data["context"] = &Context{controllerName: controllerName,
                                  actionName: actionName}
  this.Layout = "layouts/default.tpl.html"
  this.Data["Lang"] = "en-US"
}

func (this *BaseController) PrepareXsrf() {
  this.Data["xsrf_token"] = this.XsrfToken()
  this.Data["xsrf_html"] = this.XsrfFormHtml()
}

func (this *BaseController) RequireAuth() {
  username := this.Ctx.Input.Session("username")
  if username != nil {
    this.Data["loggedIn"] = true
    this.Data["username"] = username
    this.Data["permissions"] = &Permissions{Users: username == "admin",
                                            Groups: username == "admin"}
  } else {
    this.RequireAuthFailed()
  }
}

func (this *BaseController) RequireAuthFailed() {
  flash := beego.NewFlash()
  flash.Error("Not logged in!")
  flash.Store(&this.Controller)
  this.Redirect("/login", 302)
}
