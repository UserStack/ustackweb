package controllers

import (
	"github.com/UserStack/ustackweb/models"
)

type StatsController struct {
	BaseController
}

func (this *StatsController) Prepare() {
	this.PrepareXsrf()
	this.RequireAuth()
	this.RequirePermissions([]string{"list_stats"})
	this.PrepareLayout()
	this.Layout = "layouts/default.html.tpl"
}

func (this *StatsController) Index() {
	this.Data["stats"], _ = models.Stats().All()
	this.TplNames = "stats/index.html.tpl"
}
