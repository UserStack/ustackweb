package controllers

import (
	"github.com/UserStack/ustackweb/forms"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
)

type GroupsController struct {
	BaseController
}

func (this *GroupsController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *GroupsController) Index() {
	this.RequirePermissions([]string{"list_groups"})
	this.TplNames = "groups/index.html.tpl"
	groups, _ := models.Groups().AllWithoutPermissions()
	paginator := this.SetPaginator(25, int64(len(groups)))
	this.Data["groups"] = groups[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(groups))]
}

func (this *GroupsController) New() {
	this.RequirePermissions([]string{"create_groups"})
	this.TplNames = "groups/new.html.tpl"
	this.SetFormSets(&forms.NewGroup{})
}

func (this *GroupsController) Create() {
	this.RequirePermissions([]string{"create_groups"})
	this.TplNames = "groups/new.html.tpl"
	form := forms.NewGroup{}
	this.SetFormSets(&form)
	if !this.ValidFormSets(&form) {
		return
	}
	created, _, _ := models.Groups().Create(form.Name)
	if created {
		this.Redirect(beego.UrlFor("GroupsController.Index"), 302)
	} else { // backend error
		flash := beego.NewFlash()
		flash.Error("Could not create group!")
		flash.Store(&this.Controller)
		this.TplNames = "groups/new.html.tpl"
	}
}

func (this *GroupsController) Delete() {
	this.RequirePermissions([]string{"delete_groups"})
	models.Groups().Destroy(this.GetString(":id"))
	this.Redirect(beego.UrlFor("GroupsController.Index"), 302)
}
