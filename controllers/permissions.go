package controllers

import (
	"github.com/UserStack/ustackweb/forms"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
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
	paginator := pagination.SetPaginator(this.Ctx, 25, int64(len(permissions)))
	this.Data["permissions"] = permissions[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(permissions))]
}

func (this *PermissionsController) New() {
	this.RequirePermissions([]string{"create_permissions"})
	this.TplNames = "permissions/new.html.tpl"
	this.SetFormSets(&forms.NewPermission{})
}

func (this *PermissionsController) Create() {
	this.RequirePermissions([]string{"create_permissions"})
	this.TplNames = "permissions/new.html.tpl"
	form := forms.NewPermission{}
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}
	models.Permissions().Create(form.Object, form.Verb)
	this.Redirect(beego.UrlFor("PermissionsController.Index"), 302)
}

func (this *PermissionsController) Destroy() {
	this.RequirePermissions([]string{"delete_permissions"})
	models.Permissions().Destroy(this.GetString(":id"))
	this.Redirect(beego.UrlFor("PermissionsController.Index"), 302)
}
