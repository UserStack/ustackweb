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

type GroupsFilter struct {
	Type string
}

func (this GroupsFilter) AllType() bool {
	return this.Type == "all"
}

func (this GroupsFilter) NormalType() bool {
	return this.Type == "normal" || this.Type == ""
}

func (this GroupsFilter) PermissionsType() bool {
	return this.Type == "permissions"
}

func (this *GroupsController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.RequirePermissions([]string{"list_groups"})
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *GroupsController) Index() {
	this.TplNames = "groups/index.html.tpl"
	groupsFilter := GroupsFilter{Type: this.GetString(":type")}
	this.Data["groupsFilter"] = groupsFilter
	var groups []models.Group
	if groupsFilter.AllType() {
		groups, _ = models.Groups().All()
	} else if groupsFilter.NormalType() {
		groups, _ = models.Groups().AllWithoutPrefix("perm")
	} else {
		groups, _ = models.Groups().AllWithPrefix("perm")
	}
	paginator := this.SetPaginator(25, int64(len(groups)))
	this.Data["groups"] = groups[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(groups))]
}

func (this *GroupsController) New() {
	this.TplNames = "groups/new.html.tpl"
	this.SetFormSets(&forms.NewGroup{})
}

func (this *GroupsController) Create() {
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
	models.Groups().Destroy(this.GetString(":id"))
	this.Redirect(beego.UrlFor("GroupsController.Index"), 302)
}
