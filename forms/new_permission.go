package forms

import (
	"github.com/beego/i18n"
)

type NewPermission struct {
	Locale i18n.Locale `form:"-"`
	Object string      `form:"type(text)" valid:"Required"`
	Verb   string      `form:"type(text)" valid:"Required"`
}
