package forms

import (
	"github.com/UserStack/ustackweb/models"
	"github.com/beego/i18n"
	wetalkutils "github.com/beego/wetalk/modules/utils"
)

type AddUserToGroup struct {
	Locale  i18n.Locale    `form:"-"`
	GroupId int            `form:"type(select)"`
	Groups  []models.Group `form:"-"`
}

func (form *AddUserToGroup) GroupIdSelectData() [][]string {
	data := make([][]string, 0, len(form.Groups))
	for _, group := range form.Groups {
		data = append(data, []string{"group." + group.Name, wetalkutils.ToStr(group.Gid)})
	}
	return data
}
