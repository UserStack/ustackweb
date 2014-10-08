package controllers

import (
	"os"
)

type SingleSignOnController struct {
	BaseController
}

func (this *SingleSignOnController) All() {
	if this.Authenticate() {
		if this.Can["read_users"] {
			this.Ctx.Output.Header("X-Reproxy-URL", "http://"+os.Getenv("APP_1_PORT_80_TCP_ADDR")+this.Ctx.Input.Url())
			this.Ctx.Output.Header("X-Accel-Redirect", "/reproxy")
			this.Ctx.WriteString("")
		} else {
			this.DestroySession()
			this.Abort("403")
		}
	} else {
		origin := "http://" + this.Ctx.Input.Header("X-Forwarded-Host") + this.Ctx.Input.Url()
		this.Redirect("http://"+os.Getenv("USTACKWEB_PUBLIC_HOST")+"/sign_in?origin="+origin, 302)
	}
}
