package models

type PermissionCollection struct {
}

func (this *PermissionCollection) All() map[string]bool {
	return map[string]bool{
		"UserList":  false,
		"GroupList": false,
	}
}

func (this *PermissionCollection) AllByUser(name_or_uid string) (permissions map[string]bool) {
	permissions = this.All()
	groups, _ := Groups().AllByUser(name_or_uid)
	for _, group := range groups {
		if _, isPermissionGroup := permissions[group.Name]; isPermissionGroup {
			permissions[group.Name] = true
		}
	}
	return
}

func Permissions() *PermissionCollection {
	return &PermissionCollection{}
}
