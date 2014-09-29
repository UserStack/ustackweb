package routers

import (
	"github.com/UserStack/ustackweb/controllers"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func init() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("hasFormError", utils.HasFormError)

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
	beego.Router("/users/:id/enable", &controllers.UsersController{}, "get:Enable")
	beego.Router("/users/:id/disable", &controllers.UsersController{}, "get:Disable")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:Edit")
	beego.Router("/users/:id/delete", &controllers.UsersController{}, "get:Destroy")
	beego.Router("/groups", &controllers.GroupsController{})
}
