package forms

import (
	"github.com/beego/i18n"
)

type NewGroup struct {
	Locale i18n.Locale `form:"-"`
	Name   string      `form:"type(text)" valid:"Required"`
}
