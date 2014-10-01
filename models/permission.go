package models

type Permission struct {
	Internal bool   // when it belongs to user stack web
	Name     string // e.g. list_users
}

func (this *Permission) GroupName() string {
	return Permissions().GroupName(this.Name)
}
