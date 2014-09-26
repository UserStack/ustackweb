package controllers

import (
	"github.com/UserStack/ustackd/backends"
	wetalkutils "github.com/beego/wetalk/modules/utils"
	"ustackweb/utils"
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
	groups := []backends.Group{
		backends.Group{Gid: 1, Name: "Administrator"},
		backends.Group{Gid: 2, Name: "Customer"},
		backends.Group{Gid: 3, Name: "Support"},
		backends.Group{Gid: 4, Name: "Developer"},
		backends.Group{Gid: 5, Name: "Visitor"},
		backends.Group{Gid: 6, Name: "Guest"},
		backends.Group{Gid: 7, Name: "Janitor"},
		backends.Group{Gid: 8, Name: "Jupitor"}}
	paginator := wetalkutils.NewPaginator(this.Ctx.Request, 3, len(groups))
	this.Data["paginator"] = paginator
	this.Data["groups"] = groups[paginator.Offset():utils.Min(paginator.Offset()+paginator.PerPageNums, len(groups))]
}
