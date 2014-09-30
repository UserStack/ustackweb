package forms

import (
	"github.com/beego/i18n"
)

type EditPassword struct {
	Locale      i18n.Locale `form:"-"`
	OldPassword string      `form:"type(password)" valid:"Required"`
	NewPassword string      `form:"type(password)" valid:"Required;MinSize(3)"`
}

func (this *EditPassword) Placeholders() map[string]string {
	return map[string]string{
		"OldPassword": "edit_password.old_password_placeholder",
		"NewPassword": "edit_password.new_password_placeholder",
	}
}

func (this *EditPassword) Labels() map[string]string {
	return map[string]string{
		"OldPassword": "edit_password.old_password_label",
		"NewPassword": "edit_password.new_password_label",
	}
}
