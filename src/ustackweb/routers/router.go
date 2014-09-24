package routers

import (
	"ustackweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/register", &controllers.RegistrationsController{}, "get:New")
    beego.Router("/register", &controllers.RegistrationsController{}, "post:Create")
    beego.Router("/sign_in", &controllers.SessionsController{}, "get:New")
    beego.Router("/sign_in", &controllers.SessionsController{}, "post:Create")
    beego.Router("/sign_out", &controllers.SessionsController{}, "*:Destroy")
    beego.Router("/profile", &controllers.ProfileController{})
    beego.Router("/users", &controllers.UsersController{})
    beego.Router("/groups", &controllers.GroupsController{})
    beego.Router("/*", &controllers.HomeController{})
}
