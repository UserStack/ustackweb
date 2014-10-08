package controllers

import (
	"github.com/astaxie/beego"

	"github.com/UserStack/ustackweb/models"
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
	this.Data["origin"] = this.GetString("origin")
	this.TplNames = "sessions/new.html.tpl"
}

func (this *SessionsController) Create() {
	login := Login{}
	formError := this.ParseForm(&login)
	if formError == nil {
		loggedIn, uid, _ := models.Users().Login(login.Username, login.Password)
		if loggedIn {
			this.SetSession("username", login.Username)
			this.SetSession("uid", uid)
			this.RequireAuth()
			origin := this.GetString("origin")
			if origin != "" {
				this.Redirect(origin, 302)
			} else {
				this.Redirect(beego.UrlFor("HomeController.Get"), 302)
			}
			return
		}
	}
	this.RequireAuthFailed()
}

func (this *SessionsController) Destroy() {
	this.DestroySession()
	this.Redirect(beego.UrlFor("SessionsController.New"), 302)
}
