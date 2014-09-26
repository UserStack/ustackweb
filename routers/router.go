package routers

import (
	"github.com/astaxie/beego"
	"github.com/UserStack/ustackweb/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/configure", &controllers.ConfigController{}, "get:CreateAdmin")
	beego.Router("/register", &controllers.RegistrationsController{}, "get:New")
	beego.Router("/register", &controllers.RegistrationsController{}, "post:Create")
	beego.Router("/sign_in", &controllers.SessionsController{}, "get:New")
	beego.Router("/sign_in", &controllers.SessionsController{}, "post:Create")
	beego.Router("/sign_out", &controllers.SessionsController{}, "*:Destroy")
	beego.Router("/profile", &controllers.ProfileController{})
	beego.Router("/users", &controllers.UsersController{}, "get:Index")
	beego.Router("/users/new", &controllers.UsersController{}, "get:New")
	beego.Router("/users", &controllers.UsersController{}, "post:Create")
	beego.Router("/users/:id/username", &controllers.UsersController{}, "post:UpdateUsername")
	beego.Router("/users/:id/username", &controllers.UsersController{}, "get:UpdateUsername")
	beego.Router("/users/:id/password", &controllers.UsersController{}, "post:UpdatePassword")
	beego.Router("/users/:id/password", &controllers.UsersController{}, "get:UpdatePassword")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:Edit")
	beego.Router("/users/:id/delete", &controllers.UsersController{}, "get:Destroy")
	beego.Router("/groups", &controllers.GroupsController{})
}
