package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
)

type PermissionsController struct {
	BaseController
}

func (this *PermissionsController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *PermissionsController) Index() {
	this.RequirePermissions([]string{"list_permissions"})
	this.TplNames = "permissions/index.html.tpl"
	permissions := models.Permissions().All()
	paginator := this.SetPaginator(25, int64(len(permissions)))
	this.Data["permissions"] = permissions[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(permissions))]
}
