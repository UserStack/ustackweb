package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
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
	formError := this.ParseForm(&login)
	if formError == nil {
		loggedIn, _ := models.Users().Login(login.Username, login.Password)
		if loggedIn {
			this.SetSession("username", login.Username)
			this.RequireAuth()
			this.Redirect(beego.UrlFor("HomeController.Get"), 302)
			return
		}
	}
	this.RequireAuthFailed()
}

func (this *SessionsController) Destroy() {
	this.DestroySession()
	this.Redirect(beego.UrlFor("SessionsController.New"), 302)
}
