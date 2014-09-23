package routers

import (
	"ustackweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/sessions/new", &controllers.SessionsController{}, "get:New")
    beego.Router("/sessions", &controllers.SessionsController{}, "post:Create")
    beego.Router("/sessions/destroy", &controllers.SessionsController{}, "*:Destroy")
    beego.Router("/profile", &controllers.ProfileController{})
}
