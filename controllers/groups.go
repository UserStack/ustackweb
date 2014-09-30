package controllers

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
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

func (this *GroupsController) Get() {
	this.TplNames = "groups/index.html.tpl"
	groups, _ := models.Groups().All()
	paginator := this.SetPaginator(25, int64(len(groups)))
	this.Data["groups"] = groups[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(groups))]
}
