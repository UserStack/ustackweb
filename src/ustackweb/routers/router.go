package routers

import (
	"ustackweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
    beego.Router("/login", &controllers.SessionsController{}, "get:New")
    beego.Router("/login", &controllers.SessionsController{}, "post:Create")
    beego.Router("/logout", &controllers.SessionsController{}, "*:Destroy")
    beego.Router("/profile", &controllers.ProfileController{})
    beego.Router("/users", &controllers.UsersController{})
    beego.Router("/*", &controllers.HomeController{})
}
