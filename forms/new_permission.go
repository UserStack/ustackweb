package forms

import (
	"github.com/beego/i18n"
)

type NewPermission struct {
	Locale i18n.Locale `form:"-"`
	Object string      `form:"type(text)" valid:"Required"`
	Verb   string      `form:"type(text)" valid:"Required"`
}

func (this *NewPermission) Placeholders() map[string]string {
	return map[string]string{
		"Object": "new_permission.object_placeholder",
		"Verb":   "new_permission.verb_placeholder",
	}
}
