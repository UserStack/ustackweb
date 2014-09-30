package forms

import (
	"github.com/beego/i18n"
)

type EditUsername struct {
	Locale          i18n.Locale `form:"-"`
	Username        string      `form:"type(text)" valid:"Required;MinSize(3)"`
	ConfirmPassword string      `form:"type(password)" valid:"Required"`
}

func (this *EditUsername) Placeholders() map[string]string {
	return map[string]string{
		"Username":        "edit_username.username_placeholder",
		"ConfirmPassword": "edit_username.confirm_password_placeholder",
	}
}

func (this *EditUsername) Labels() map[string]string {
	return map[string]string{
		"ConfirmPassword": "edit_username.confirm_password_label",
	}
}
