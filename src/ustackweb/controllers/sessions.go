package controllers

import (
	"github.com/astaxie/beego"
	"ustackweb/models"
)

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
	this.TplNames = "sessions/new.html.tpl"
}

func (this *SessionsController) Create() {
	login := Login{}
	err := this.ParseForm(&login)
	if err == nil && models.Users().Login(login.Username, login.Password) {
		this.SetSession("username", login.Username)
		this.RequireAuth()
		this.Redirect(beego.UrlFor("ProfileController.Get"), 302)
	} else {
		this.RequireAuthFailed()
	}
}

func (this *SessionsController) Destroy() {
	this.DestroySession()
	this.Redirect(beego.UrlFor("SessionsController.New"), 302)
}
