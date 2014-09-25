package routers

import (
	"github.com/astaxie/beego"
	"ustackweb/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegistrationsController{}, "get:New")
	beego.Router("/register", &controllers.RegistrationsController{}, "post:Create")
	beego.Router("/sign_in", &controllers.SessionsController{}, "get:New")
	beego.Router("/sign_in", &controllers.SessionsController{}, "post:Create")
	beego.Router("/sign_out", &controllers.SessionsController{}, "*:Destroy")
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/users", &controllers.UsersController{}, "get:Index")
	beego.Router("/users/:id/update", &controllers.UsersController{}, "post:Update")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:Edit")
	beego.Router("/groups", &controllers.GroupsController{})
	beego.Router("/*", &controllers.HomeController{})
}
