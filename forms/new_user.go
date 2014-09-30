package forms

import (
	"github.com/beego/i18n"
)

type NewUser struct {
	Locale   i18n.Locale `form:"-"`
	Username string      `form:"type(text)" valid:"Required;MinSize(3)"`
	Password string      `form:"type(password)" valid:"Required;MinSize(3)"`
}

func (this *NewUser) Placeholders() map[string]string {
	return map[string]string{
		"Username": "new_user.username_placeholder",
		"Password": "new_user.password_placeholder",
	}
}
