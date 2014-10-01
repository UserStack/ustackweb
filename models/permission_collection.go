package models

type PermissionCollection struct {
}

type Permission struct {
	Key string // UserList
	Raw string // perm.user.list
}

func (this *PermissionCollection) All() map[string]bool {
	return map[string]bool{
		"ustack-perm-user-list":   false,
		"ustack-perm-user-read":   false,
		"ustack-perm-user-write":  false,
		"ustack-perm-group-list":  false,
		"ustack-perm-group-read":  false,
		"ustack-perm-group-write": false,
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
