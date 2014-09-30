package forms

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/beego/i18n"
)

type NewUser struct {
	Locale   i18n.Locale    `form:"-"`
	Username string         `form:"type(text)" valid:"Required;MinSize(3)"`
	Password string         `form:"type(password)" valid:"Required;MinSize(3)"`
	Group    models.Group   `form:"type(select);attr(multiple)"`
	Groups   []models.Group `form:"-"`
}

func (this *NewUser) Placeholders() map[string]string {
	return map[string]string{
		"Username": "new_user.username_placeholder",
		"Password": "new_user.password_placeholder",
	}
}

func (form *NewUser) GroupSelectData() [][]string {
	data := make([][]string, 0, len(form.Groups))
	for _, group := range form.Groups {
		data = append(data, []string{"group." + group.Name, string(group.Gid)})
	}
	return data
}
