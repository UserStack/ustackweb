package models

type Permission struct {
	Name string // e.g. list_users
}

func (this *Permission) GroupName() string {
	return Permissions().GroupName(this.Name)
}
