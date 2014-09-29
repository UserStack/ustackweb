package forms

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego/validation"
)

type EditUsername struct {
	XsrfHtml         string
	ValidationErrors []*validation.ValidationError
	User             *models.User
}

type EditPassword struct {
	XsrfHtml         string
	ValidationErrors []*validation.ValidationError
	User             *models.User
	Password         string
	OldPassword      string
}
