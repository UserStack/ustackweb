package routers

import (
	"github.com/UserStack/ustackweb/controllers"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func init() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("compare", beego.Compare)
	beego.AddFuncMap("hasFormError", utils.HasFormError)

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/install", &controllers.InstallController{}, "*:Index")
	beego.Router("/install/create_root_user", &controllers.InstallController{}, "*:CreateRootUser")
	beego.Router("/install/create_permissions", &controllers.InstallController{}, "*:CreatePermissions")
	beego.Router("/install/assign_permissions", &controllers.InstallController{}, "*:AssignPermissions")
	beego.Router("/install/drop_database", &controllers.InstallController{}, "*:DropDatabase")
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
	beego.Router("/users/:id/username", &controllers.UsersController{}, "get:EditUsername")
	beego.Router("/users/:id/password", &controllers.UsersController{}, "post:UpdatePassword")
	beego.Router("/users/:id/password", &controllers.UsersController{}, "get:EditPassword")
	beego.Router("/users/:id/enable", &controllers.UsersController{}, "get:Enable")
	beego.Router("/users/:id/disable", &controllers.UsersController{}, "get:Disable")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:Edit")
	beego.Router("/users/:id/edit_groups", &controllers.UsersController{}, "get:EditGroups")
	beego.Router("/users/:id/groups/:groupId/add", &controllers.UsersController{}, "*:AddUserToGroup")
	beego.Router("/users/:id/groups/:groupId/remove", &controllers.UsersController{}, "*:RemoveUserFromGroup")
	beego.Router("/users/:id/delete", &controllers.UsersController{}, "get:Destroy")
	beego.Router("/groups", &controllers.GroupsController{}, "get:Index")
	beego.Router("/groups/new", &controllers.GroupsController{}, "get:New")
	beego.Router("/groups/:id/delete", &controllers.GroupsController{}, "get:Delete")
	beego.Router("/groups", &controllers.GroupsController{}, "post:Create")
	beego.Router("/permissions", &controllers.PermissionsController{}, "get:Index")
	beego.Router("/permissions/new", &controllers.PermissionsController{}, "get:New")
	beego.Router("/permissions", &controllers.PermissionsController{}, "post:Create")
	beego.Router("/stats", &controllers.StatsController{}, "get:Index")
}
