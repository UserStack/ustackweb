package controllers

import (
	"fmt"
	"github.com/UserStack/ustackweb/models"
	"github.com/UserStack/ustackweb/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	wetalkutils "github.com/beego/wetalk/modules/utils"
	"reflect"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func (this *BaseController) PrepareLayout() {
	this.Data["context"] = utils.NewContext(this.GetControllerAndAction())
	this.Layout = "layouts/default.html.tpl"
	this.Data["Lang"] = "en-US"
	this.Locale.Lang = "en-US"
}

func (this *BaseController) PrepareXsrf() {
	this.Data["xsrf_token"] = this.XsrfToken()
	this.Data["xsrf_html"] = this.XsrfFormHtml()
}

func (this *BaseController) RequireAuth() {
	username := this.Ctx.Input.Session("username")
	if username != nil {
		this.Data["loggedIn"] = true
		this.Data["username"] = username
		this.Data["can"] = models.Permissions().Abilities(fmt.Sprintf("%s", username))
	} else {
		this.RequireAuthFailed()
	}
}

func (this *BaseController) RequireAuthFailed() {
	flash := beego.NewFlash()
	flash.Error("Not logged in!")
	flash.Store(&this.Controller)
	this.Redirect("/sign_in", 302)
}

func (this *BaseController) SetPaginator(per int, nums int64) *wetalkutils.Paginator {
	p := wetalkutils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) SetFormSets(form interface{}, names ...string) *wetalkutils.FormSets {
	return this.setFormSets(form, nil, names...)
}

func (this *BaseController) setFormSets(form interface{}, errs map[string]*validation.ValidationError, names ...string) *wetalkutils.FormSets {
	formSets := wetalkutils.NewFormSets(form, errs, this.Locale)
	name := reflect.ValueOf(form).Elem().Type().Name()
	if len(names) > 0 {
		name = names[0]
	}
	name += "FormSets"
	this.Data[name] = formSets

	return formSets
}

func (this *BaseController) validForm(form interface{}, names ...string) (bool, map[string]*validation.ValidationError) {
	// parse request params to form ptr struct
	wetalkutils.ParseForm(form, this.Input())

	// Put data back in case users input invalid data for any section.
	name := reflect.ValueOf(form).Elem().Type().Name()
	if len(names) > 0 {
		name = names[0]
	}
	this.Data[name] = form

	errName := name + "Error"

	// Verify basic input.
	valid := validation.Validation{}
	if ok, _ := valid.Valid(form); !ok {
		errs := valid.ErrorMap()
		this.Data[errName] = &valid
		return false, errs
	}
	return true, nil
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidForm(form interface{}, names ...string) bool {
	valid, _ := this.validForm(form, names...)
	return valid
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidFormSets(form interface{}, names ...string) bool {
	valid, errs := this.validForm(form, names...)
	this.setFormSets(form, errs, names...)
	return valid
}
